package models

import (
	"github.com/go-playground/validator/v10"
	"hackathon-basic-backend/database"
	"hackathon-basic-backend/helpers"
	"net/http"
)

type Transporter struct {
	Id      int    `json:"id"`
	LicenseNumber    string `json:"license_number" validate:"required"`
	LicenseType    string `json:"license_type" validate:"required"`
	TruckType    string `json:"truck_type" validate:"required"`
	ProductionYear    string `json:"production_year" validate:"required"`
	Status    bool `json:"status" validate:"required"`
}

func FetchAllTransporters() (Response, error) {
	var obj Transporter
	var arrayObj []Transporter
	var res Response

	con := database.GetConnection()

	sqlStatement := "SELECT * FROM transporters"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helpers.OutputPanicError(err)

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.LicenseNumber, &obj.LicenseType, &obj.TruckType, &obj.ProductionYear, &obj.TruckType)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayObj

	return res, nil
}

func StoreTransporter(licenseNumber, licenseType, truckType string) (Response, error) {
	var res Response

	v := validator.New()
	transporterStruct := Transporter{
		LicenseNumber:    licenseNumber,
		LicenseType: licenseType,
		TruckType:   truckType,
	}
	// validation input
	err := v.Struct(transporterStruct)
	if err != nil {
		return res, err
	}

	con := database.GetConnection()

	sqlStatement := "INSERT transporters (license_number, license_type, truck_type) VALUES (?,?,?)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(licenseNumber, licenseType, truckType)
	if err != nil {
		return res, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully created"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateTransporter(id int, licenseNumber, licenseType, truckType string) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := "UPDATE transporters set license_number = ?, license_type = ?, truck_type = ? WHERE id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(licenseNumber, licenseType, truckType, id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully updated"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteTransporter(id int) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := "DELETE FROM transporters WHERE id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully deleted"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
