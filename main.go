package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var db *sql.DB

var cfg = mysql.Config{
	User:      "user",
	Passwd:    "password",
	Net:       "tcp",
	Addr:      "localhost:3306",
	DBName:    "db",
	ParseTime: true,
}

func main() {
	db, _ = sql.Open("mysql", cfg.FormatDSN())
	defer db.Close()

	allowOrigins := handlers.AllowedOrigins([]string{"*"})
	allowMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	allowHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passenger/{id}", filterPassenger).Methods(http.MethodGet, http.MethodPut)
	router.HandleFunc("/api/v1/passenger", passenger).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/api/v1/driver/is_available/{id}", isAvailableDriver).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/driver/{id}", filterDriver).Methods(http.MethodGet, http.MethodPut)
	router.HandleFunc("/api/v1/driver", driver).Methods(http.MethodGet, http.MethodPost)

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowOrigins, allowMethods, allowHeaders)(router)))
}
func passenger(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		pList, err := getPassenger()
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(pList)
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}
		fmt.Println(pList)

	case http.MethodPost:
		var p Passenger
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &p); ok == nil {
				err := insertPassenger(p)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "%s inserted", p.First_Name)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				}
			}
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}
func filterPassenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodGet:

		pList, err := getPassengerFilterId(&id)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(pList)
			fmt.Println(pList)

			fmt.Println(string(out))
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}

	case http.MethodPut:
		var p Passenger
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &p); ok == nil {
				err := updatePassenger(id, p)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "%s updated", p.First_Name)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				}
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}

func driver(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		pList, err := getDriver()
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(pList)
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}
		fmt.Println(pList)

	case http.MethodPost:
		var d Driver
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &d); ok == nil {
				err := insertDriver(d)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "%s inserted", d.First_Name)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				}
			}
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}
func isAvailableDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodPut:
		var d Driver
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &d); ok == nil {
				err := updateDriverAvailability(id, d)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "%s updated", d.First_Name)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				}
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}

func filterDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodGet:

		pList, err := getDriverFilterId(&id)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(pList)
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}
		fmt.Println(pList)

	case http.MethodPut:
		var d Driver
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &d); ok == nil {
				err := updateDriver(id, d)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "%s updated", d.First_Name)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				}
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}
