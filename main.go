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
	"github.com/gorilla/mux"
)

var db *sql.DB

var cfg = mysql.Config{
	User:   "user",
	Passwd: "password",
	Net:    "tcp",
	Addr:   "localhost:3306",
	DBName: "db",
}

type Passanger struct {
	Passanger_Id int    `json:"passanger_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
}

func main() {
	db, _ = sql.Open("mysql", cfg.FormatDSN())

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passanger/{id}", filterPassanger)
	router.HandleFunc("/api/v1/passanger", passanger)

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
func passanger(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		pList, err := getPassanger(nil)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(pList)
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

		pList, err := getPassanger(&id)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			out, _ := json.Marshal(pList)
			fmt.Fprintf(w, string(out))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}
		fmt.Println(pList)

	case http.MethodPut:
		var p Passanger
		if byteBody, ok := ioutil.ReadAll(r.Body); ok == nil {
			if ok := json.Unmarshal(byteBody, &p); ok == nil {
				err := updatePassanger(id, p)
				if err == nil {
					w.WriteHeader(http.StatusAccepted)
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

func getPassanger(id *int) ([]Passanger, error) {
	var pList []Passanger
	var rows *sql.Rows
	var err error

	if id == nil {
		rows, err = db.Query("SELECT * FROM passanger")
	} else {
		rows, err = db.Query("SELECT * FROM passanger WHERE passanger_id = ?", *id)

	}

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Passanger
		if err := rows.Scan(&p.Passanger_Id, &p.First_Name, &p.Last_Name, &p.Mobile_No, &p.Email); err != nil {
			return nil, err
		}
		pList = append(pList, p)
	}
	return pList, nil
}
func insertPassanger(p Passanger) error {
	_, err := db.Query("INSERT INTO passanger(passanger_id, first_name, last_name, mobile_no, email) VALUES (?, ?, ?, ?, ?)", p.Passanger_Id, p.First_Name, p.Last_Name, p.Mobile_No, p.Email)
	return err
}

func updatePassanger(id int, p Passanger) error {
	_, err := db.Query("UPDATE passanger SET first_name = ?, last_name = ?, mobile_no = ?, email = ? WHERE passanger_id = ?", p.First_Name, p.Last_Name, p.Mobile_No, p.Email, id)
	return err
}
