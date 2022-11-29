package main

import "database/sql"

func getPassanger() ([]Passanger, error) {
	var pList []Passanger
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
	var pList []Passanger
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
	var dList []Driver
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM driver")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Driver
		if err := rows.Scan(&p.Driver_Id, &p.First_Name, &p.Last_Name, &p.Mobile_No, &p.Email, &p.Id_No, &p.Car_No); err != nil {
			return nil, err
		}
		dList = append(dList, p)
	}
	return dList, nil
}
func getDriverFilterId(id *int) ([]Driver, error) {
	var dList []Driver
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM driver WHERE driver_id = ?", *id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Driver
		if err := rows.Scan(&p.Driver_Id, &p.First_Name, &p.Last_Name, &p.Mobile_No, &p.Email, &p.Id_No, &p.Car_No); err != nil {
			return nil, err
		}
		dList = append(dList, p)
	}
	return dList, nil
}
func insertDriver(d Driver) error {
	_, err := db.Query("INSERT INTO driver(driver_id, first_name, last_name, mobile_no, email, id_no, car_no) VALUES (?, ?, ?, ?, ?, ?, ?)", d.Driver_Id, d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Id_No, d.Car_No)
	return err
}

func updateDriver(id int, d Driver) error {
	_, err := db.Query("UPDATE driver SET first_name = ?, last_name = ?, mobile_no = ?, email = ?, id_no = ?, car_no = ? WHERE driver_id = ?", d.First_Name, d.Last_Name, d.Mobile_No, d.Email, d.Id_No, d.Car_No, id)
	return err
}