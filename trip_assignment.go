package main

import "database/sql"

type Trip_Assignment struct {
	Trip_Id   int    `json:"trip_id"`
	Driver_Id int    `json:"driver_id"`
	Status    string `json:"status"`
}

type Trip_Assignment_With_Passanger_Trip struct {
	Trip_Id      int            `json:"trip_id"`
	Driver_Id    sql.NullInt32  `json:"driver_id"`
	Status       sql.NullString `json:"status"`
	Passanger_Id int            `json:"passanger_id"`
	First_Name   sql.NullString `json:"first_name"`
	Last_Name    sql.NullString `json:"last_name"`
	Mobile_No    sql.NullString `json:"mobile_no"`
	Email        sql.NullString `json:"email"`
	Pick_Up      string         `json:"pick_up"`
	Drop_Off     string         `json:"drop_off"`
	Start        sql.NullTime   `json:"start"`
	End          sql.NullTime   `json:"end"`
	Car_No       sql.NullString `json:"car_no"`
}

type Trip_Assignment_With_Driver_Trip struct {
	Trip_Id      int          `json:"trip_id"`
	Driver_Id    int          `json:"driver_id"`
	Status       string       `json:"status"`
	Passanger_Id int          `json:"passanger_id"`
	First_Name   string       `json:"first_name"`
	Last_Name    string       `json:"last_name"`
	Mobile_No    string       `json:"mobile_no"`
	Email        string       `json:"email"`
	Pick_Up      string       `json:"pick_up"`
	Drop_Off     string       `json:"drop_off"`
	Start        sql.NullTime `json:"start"`
	End          sql.NullTime `json:"end"`
}
