package main

import "database/sql"

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
	_, err := db.Query("INSERT INTO driver(driver_id, first_name, last_name, mobile_no, email, id_no, car_no) VALUES (?, ?, ?, ?, ?, ?, ?)", d.Driver_Id, d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Id_No, d.Car_No, d.Is_Available)
	return err
}

func updateDriver(id int, d Driver) error {
	_, err := db.Query("UPDATE driver SET first_name = ?, last_name = ?, mobile_no = ?, email = ?, id_no = ?, car_no = ?, is_available = ? WHERE driver_id = ?", d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Id_No, d.Car_No, d.Is_Available, id)
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
