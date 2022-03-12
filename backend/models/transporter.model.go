package models

import (
	"net/http"
)

type Transporter struct {
	Id             int    `json:"id"`
	LicenseNumber  string `json:"license_number" validate:"required"`
	LicenseType    string `json:"license_type" validate:"required"`
	TruckType      string `json:"truck_type" validate:"required"`
	ProductionYear string `json:"production_year" validate:"required"`
	Status         bool   `json:"status" validate:"required"`
}

func FetchAllTransporters() (Response, error) {
	var res Response

	transporters := []Transporter{}

	DB.Find(&transporters)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = transporters

	return res, nil
}

func StoreTransporter(licenseNumber, licenseType, truckType string) (Response, error) {
	var res Response
	return res, nil
}

func UpdateTransporter(id int, licenseNumber, licenseType, truckType string) (Response, error) {
	var res Response
	return res, nil
}

func DeleteTransporter(id int) (Response, error) {
	var res Response

	return res, nil
}
