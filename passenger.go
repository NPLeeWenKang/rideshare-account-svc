package main

import "database/sql"

type Passenger struct {
	Passenger_Id int    `json:"passenger_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
}

func getPassenger() ([]Passenger, error) {
	dList := make([]Passenger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM passenger")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data Passenger
		if err := rows.Scan(&data.Passenger_Id, &data.First_Name, &data.Last_Name, &data.Mobile_No, &data.Email); err != nil {
			return nil, err
		}
		dList = append(dList, data)
	}
	return dList, nil
}
func getPassengerFilterId(id *int) ([]Passenger, error) {
	dList := make([]Passenger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM passenger WHERE passenger_id = ?", *id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data Passenger
		if err := rows.Scan(&data.Passenger_Id, &data.First_Name, &data.Last_Name, &data.Mobile_No, &data.Email); err != nil {
			return nil, err
		}
		dList = append(dList, data)
	}
	return dList, nil
}
func insertPassenger(p Passenger) error {
	_, err := db.Query("INSERT INTO passenger(passenger_id, first_name, last_name, mobile_no, email) VALUES (?, ?, ?, ?, ?)", p.Passenger_Id, p.First_Name, p.Last_Name, p.Mobile_No, p.Email)
	return err
}

func updatePassenger(id int, p Passenger) error {
	_, err := db.Query("UPDATE passenger SET first_name = ?, last_name = ?, mobile_no = ?, email = ? WHERE passenger_id = ?", p.First_Name, p.Last_Name, p.Mobile_No, p.Email, id)
	return err
}
