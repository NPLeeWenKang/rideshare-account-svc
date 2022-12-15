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
	router.HandleFunc("/api/v1/passanger/current_assignment/{id}", currentAssignmentPassanger).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/passanger/{id}", filterPassanger).Methods(http.MethodGet, http.MethodPut)
	router.HandleFunc("/api/v1/passanger", passanger).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/api/v1/driver/is_available/{id}", isAvailableDriver).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/driver/current_assignment/{id}", currentAssignmentDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/driver/{id}", filterDriver).Methods(http.MethodGet, http.MethodPut)
	router.HandleFunc("/api/v1/driver", driver).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/api/v1/trip/{id}", filterTrip).Methods(http.MethodGet, http.MethodPut)
	router.HandleFunc("/api/v1/trip", trip).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/api/v1/trip_assignment", tripAssignment).Methods(http.MethodPut)

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowOrigins, allowMethods, allowHeaders)(router)))
}
func passanger(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		pList, err := getPassanger()
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
		var p Passanger
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &p); ok == nil {
				err := insertPassanger(p)
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
func filterPassanger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodGet:

		pList, err := getPassangerFilterId(&id)
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
		var p Passanger
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &p); ok == nil {
				err := updatePassanger(id, p)
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

func trip(w http.ResponseWriter, r *http.Request) {
	querystringmap := r.URL.Query()
	passangerId := querystringmap["passanger_id"]

	switch r.Method {
	case http.MethodGet:
		if len(passangerId) >= 1 {
			tList, err := getTripFilterPassangerId(passangerId[0])
			if err == nil {
				w.WriteHeader(http.StatusAccepted)
				out, _ := json.Marshal(tList)
				w.Header().Set("Content-type", "application/json")
				fmt.Fprintf(w, string(out))
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			}
		} else {
			tList, err := getTrip()
			if err == nil {
				w.WriteHeader(http.StatusAccepted)
				out, _ := json.Marshal(tList)
				w.Header().Set("Content-type", "application/json")
				fmt.Fprintf(w, string(out))
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			}
		}

	case http.MethodPost:
		var t Trip
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &t); ok == nil {
				err := insertTrip(t)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "Inserted Trip Id %d", t.Trip_Id)
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
func filterTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodGet:

		tList, err := getTripFilterId(&id)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(tList)
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}

	case http.MethodPut:
		var t Trip
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &t); ok == nil {
				err := updateTrip(id, t)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "Updated Trip Id %d", t.Trip_Id)
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

func currentAssignmentPassanger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodGet:
		tList, err := getCurrentTripAssignmentFilterPassangerId(id)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(tList)
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}

func currentAssignmentDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No ID")
	}
	id, _ := strconv.Atoi(params["id"])

	switch r.Method {
	case http.MethodGet:
		tList, err := getCurrentTripAssignmentFilterDriverId(id)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(tList)
			w.Header().Set("Content-type", "application/json")
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Error")
	}
}

func tripAssignment(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
		var ta Trip_Assignment
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &ta); ok == nil {
				var err error
				if ta.Status == "ACCEPTED" {
					err = updateTripAssignmentAndTripStart(ta)
				} else if ta.Status == "DRIVING" {
					err = updateTripAssignmentAndTripEnd(ta)
				} else {
					err = updateTripAssignment(ta)
				}
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
					w.Header().Set("Content-type", "application/json")
					fmt.Fprintf(w, "trip_id %d driver_id %d updated", ta.Trip_Id, ta.Driver_Id)
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
