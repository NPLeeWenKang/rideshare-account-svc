package main

import "database/sql"

type Passanger struct {
	Passanger_Id int    `json:"passanger_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
}

func getPassanger() ([]Passanger, error) {
	dList := make([]Passanger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM passanger")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data Passanger
		if err := rows.Scan(&data.Passanger_Id, &data.First_Name, &data.Last_Name, &data.Mobile_No, &data.Email); err != nil {
			return nil, err
		}
		dList = append(dList, data)
	}
	return dList, nil
}
func getPassangerFilterId(id *int) ([]Passanger, error) {
	dList := make([]Passanger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM passanger WHERE passanger_id = ?", *id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data Passanger
		if err := rows.Scan(&data.Passanger_Id, &data.First_Name, &data.Last_Name, &data.Mobile_No, &data.Email); err != nil {
			return nil, err
		}
		dList = append(dList, data)
	}
	return dList, nil
}
func insertPassanger(p Passanger) error {
	_, err := db.Query("INSERT INTO passanger(passanger_id, first_name, last_name, mobile_no, email) VALUES (?, ?, ?, ?, ?)", p.Passanger_Id, p.First_Name, p.Last_Name, p.Mobile_No, p.Email)
	return err
}

func updatePassanger(id int, p Passanger) error {
	_, err := db.Query("UPDATE passanger SET first_name = ?, last_name = ?, mobile_no = ?, email = ? WHERE passanger_id = ?", p.First_Name, p.Last_Name, p.Mobile_No, p.Email, id)
	return err
}
