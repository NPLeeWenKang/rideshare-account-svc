package main

import (
	"database/sql"
	"fmt"
)

type Driver struct {
	Driver_Id    int    `json:"driver_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
	Id_No        string `json:"id_no"`
	Car_No       string `json:"car_no"`
	Is_Available bool   `json:"is_available"`
}

func getDriver() ([]Driver, error) {
	dList := make([]Driver, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM driver")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data Driver
		if err := rows.Scan(&data.Driver_Id, &data.First_Name, &data.Last_Name, &data.Mobile_No, &data.Email, &data.Id_No, &data.Car_No, &data.Is_Available); err != nil {
			return nil, err
		}
		dList = append(dList, data)
	}
	return dList, nil
}
func getDriverFilterId(id *int) ([]Driver, error) {
	dList := make([]Driver, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM driver WHERE driver_id = ?", *id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data Driver
		if err := rows.Scan(&data.Driver_Id, &data.First_Name, &data.Last_Name, &data.Mobile_No, &data.Email, &data.Id_No, &data.Car_No, &data.Is_Available); err != nil {
			return nil, err
		}
		dList = append(dList, data)
	}
	return dList, nil
}
func insertDriver(d Driver) error {
	fmt.Println(d)
	_, err := db.Query("INSERT INTO driver(first_name, last_name, mobile_no, email, id_no, car_no, is_available) VALUES (?, ?, ?, ?, ?, ?, false)", d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Id_No, d.Car_No)
	fmt.Println(err)
	return err
}

func updateDriver(id int, d Driver) error {
	_, err := db.Query("UPDATE driver SET first_name = ?, last_name = ?, mobile_no = ?, email = ?, car_no = ? WHERE driver_id = ?", d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Car_No, id)
	return err
}
func updateDriverAvailability(id int, d Driver) error {
	fmt.Println(id, d.Is_Available)
	_, err := db.Query("UPDATE driver SET is_available = ? WHERE driver_id = ?", d.Is_Available, id)
	return err
}
