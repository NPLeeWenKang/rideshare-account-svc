package main

import (
	"database/sql"
	"fmt"
)

func getPassanger() ([]Passanger, error) {
	pList := make([]Passanger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM passanger")

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
func getPassangerFilterId(id *int) ([]Passanger, error) {
	pList := make([]Passanger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM passanger WHERE passanger_id = ?", *id)

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

func getDriver() ([]Driver, error) {
	dList := make([]Driver, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM driver")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Driver
		if err := rows.Scan(&p.Driver_Id, &p.First_Name, &p.Last_Name, &p.Mobile_No, &p.Email, &p.Id_No, &p.Car_No, &p.Is_Available); err != nil {
			return nil, err
		}
		dList = append(dList, p)
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
		var p Driver
		if err := rows.Scan(&p.Driver_Id, &p.First_Name, &p.Last_Name, &p.Mobile_No, &p.Email, &p.Id_No, &p.Car_No, &p.Is_Available); err != nil {
			return nil, err
		}
		dList = append(dList, p)
	}
	return dList, nil
}
func insertDriver(d Driver) error {
	_, err := db.Query("INSERT INTO driver(driver_id, first_name, last_name, mobile_no, email, id_no, car_no) VALUES (?, ?, ?, ?, ?, ?, false)", d.Driver_Id, d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Id_No, d.Car_No)
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

func getTrip() ([]Trip, error) {
	tList := make([]Trip, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM trip")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Trip
		if err := rows.Scan(&t.Trip_Id, &t.Passanger_Id, &t.Pick_Up, &t.Drop_Off, &t.Start, &t.End); err != nil {
			return nil, err
		}
		tList = append(tList, t)
	}
	return tList, nil
}

func getTripFilterId(id *int) ([]Trip, error) {
	tList := make([]Trip, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM trip WHERE trip_id = ?", *id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Trip
		if err := rows.Scan(&t.Trip_Id, &t.Passanger_Id, &t.Pick_Up, &t.Drop_Off, &t.Start, &t.End); err != nil {
			return nil, err
		}
		tList = append(tList, t)
	}
	return tList, nil
}

func getTripFilterPassangerId(passangerId string) ([]Trip_Filter_Passanger, error) {
	tList := make([]Trip_Filter_Passanger, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("WITH latest_assignment AS ( SELECT ta1.* FROM trip_assignment ta1 LEFT JOIN trip_assignment ta2 ON ta1.trip_id = ta2.trip_id AND ta1.assign_datetime < ta2.assign_datetime WHERE ta2.trip_id is NULL ), latest_trip AS ( SELECT t.trip_id, la.status FROM trip t LEFT JOIN latest_assignment la ON t.trip_id = la.trip_id ) SELECT t.*, lt.status FROM passanger p INNER JOIN trip t ON p.passanger_id = t.passanger_id INNER JOIN latest_trip lt ON t.trip_id = lt.trip_id WHERE p.passanger_id = ?", passangerId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Trip_Filter_Passanger
		if err := rows.Scan(&t.Trip_Id, &t.Passanger_Id, &t.Pick_Up, &t.Drop_Off, &t.Start, &t.End, &t.Status); err != nil {
			return nil, err
		}
		tList = append(tList, t)
	}
	return tList, nil
}

func insertTrip(t Trip) error {
	_, err := db.Query("INSERT INTO trip(trip_id, passanger_id, pick_up, drop_off, start, end) VALUES (?, ?, ?, ?, ?, ?)", t.Trip_Id, t.Passanger_Id, t.Pick_Up, t.Drop_Off, t.Start, t.End)
	return err
}

func updateTrip(id int, t Trip) error {
	_, err := db.Query("UPDATE trip SET passanger_id = ?, pick_up = ?, drop_off = ?, start = ?, end = ? WHERE trip_id = ?", t.Passanger_Id, t.Pick_Up, t.Drop_Off, t.Start, t.End, id)
	return err
}

func getCurrentTripAssignmentFilterDriverId(driverId int) ([]Trip_Assignment_With_Driver_Trip, error) {
	tList := make([]Trip_Assignment_With_Driver_Trip, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("WITH latest_assignment AS ( SELECT ta1.* FROM trip_assignment ta1 LEFT JOIN trip_assignment ta2 ON ta1.trip_id = ta2.trip_id AND ta1.assign_datetime < ta2.assign_datetime WHERE ta2.trip_id is NULL AND ta1.status != 'DONE' AND ta1.status != 'REJECTED' ) SELECT d.driver_id, t.*, la.status, p.first_name, p.last_name, p.mobile_no, p.email  FROM latest_assignment la INNER JOIN trip t ON la.trip_id = t.trip_id INNER JOIN passanger p ON t.passanger_id = p.passanger_id INNER JOIN driver d ON d.driver_id = la.driver_id WHERE d.driver_id = ?", driverId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Trip_Assignment_With_Driver_Trip
		if err := rows.Scan(&t.Driver_Id, &t.Trip_Id, &t.Passanger_Id, &t.Pick_Up, &t.Drop_Off, &t.Start, &t.End, &t.Status, &t.First_Name, &t.Last_Name, &t.Mobile_No, &t.Email); err != nil {
			return nil, err
		}
		tList = append(tList, t)
	}
	return tList, nil
}

func getCurrentTripAssignmentFilterPassangerId(passangerId int) ([]Trip_Assignment_With_Passanger_Trip, error) {
	tList := make([]Trip_Assignment_With_Passanger_Trip, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("WITH latest_assignment AS ( SELECT ta1.* FROM trip_assignment ta1 LEFT JOIN trip_assignment ta2 ON ta1.trip_id = ta2.trip_id AND ta1.assign_datetime < ta2.assign_datetime WHERE ta2.trip_id is NULL AND ta1.status != 'DONE' AND ta1.status != 'REJECTED' ) SELECT d.driver_id, t.*, la.status, d.first_name, d.last_name, d.mobile_no, d.email, d.car_no  FROM latest_assignment la INNER JOIN trip t ON la.trip_id = t.trip_id INNER JOIN passanger p ON t.passanger_id = p.passanger_id INNER JOIN driver d ON d.driver_id = la.driver_id WHERE p.passanger_id = ?", passangerId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Trip_Assignment_With_Passanger_Trip
		if err := rows.Scan(&t.Driver_Id, &t.Trip_Id, &t.Passanger_Id, &t.Pick_Up, &t.Drop_Off, &t.Start, &t.End, &t.Status, &t.First_Name, &t.Last_Name, &t.Mobile_No, &t.Email, &t.Car_No); err != nil {
			return nil, err
		}
		tList = append(tList, t)
	}
	return tList, nil
}

func updateTripAssignment(t Trip_Assignment) error {
	fmt.Println(t)
	_, err := db.Query("UPDATE trip_assignment SET status = ? WHERE trip_id = ? AND driver_id = ?", t.Status, t.Trip_Id, t.Driver_Id)
	return err
}

func updateTripAssignmentAndTripStart(t Trip_Assignment) error {
	fmt.Println(t)
	_, err := db.Query("START TRANSACTION; UPDATE trip_assignment SET status = ? WHERE trip_id = ? AND driver_id = ?; UPDATE trip SET start = CURRENT_TIMESTAMP WHERE trip_id = ?; COMMIT;", t.Status, t.Trip_Id, t.Driver_Id, t.Trip_Id)
	return err
}

func updateTripAssignmentAndTripEnd(t Trip_Assignment) error {
	fmt.Println(t)
	_, err := db.Query("START TRANSACTION; UPDATE trip_assignment SET status = ? WHERE trip_id = ? AND driver_id = ?; UPDATE trip SET end = CURRENT_TIMESTAMP WHERE trip_id = ?; COMMIT;", t.Status, t.Trip_Id, t.Driver_Id, t.Trip_Id)
	return err
}
