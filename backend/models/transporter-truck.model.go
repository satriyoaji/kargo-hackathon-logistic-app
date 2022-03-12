package models

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Truck struct {
	Id             int    `json:"id"`
	LicenseNumber  string `json:"license_number" validate:"required"`
	TruckType      string `json:"truck_type" validate:"required"`
	PlateType      string `json:"plate_type" validate:"required"`
	ProductionYear string `json:"production_year" validate:""`
	STNK           string `json:"stnk" validate:""`
	KIR            string `json:"kir" validate:""`
	Status         bool   `json:"status" validate:"" gorm:"not null"`
}

func FetchAllTrucks() (Response, error) {
	var res Response

	trucks := []Truck{}

	DB.Find(&trucks)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = trucks

	return res, nil
}

func FetchByIdTruck(id int) (Response, error) {
	var res Response

	truck := Truck{}

	if err := DB.Where("id = ?", id).First(&truck).Error; err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "the truck not found !"
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = truck

	return res, nil
}

func StoreTruck(licenseNumber, truckType, plateType, productionYear string) (Response, error) {
	var res Response

	v := validator.New()
	truckStruct := Truck{
		LicenseNumber:  licenseNumber,
		TruckType:      truckType,
		PlateType:      plateType,
		ProductionYear: productionYear,
		Status:         true,
	}
	// validation input
	err := v.Struct(truckStruct)
	if err != nil {
		return res, err
	}

	DB.Create(&truckStruct)

	res.Status = http.StatusOK
	res.Message = "new truck successfully created"
	res.Data = truckStruct
	return res, nil
}

func UpdateTruck(id int, licenseNumber, truckType, plateType, productionYear string, status bool) (Response, error) {
	var res Response

	var truck Truck
	if err := DB.Where("id = ?", id).First(&truck).Error; err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "the truck which will be updated not found !"
		return res, err
	}

	v := validator.New()
	truckStruct := Truck{
		Id:             id,
		LicenseNumber:  licenseNumber,
		TruckType:      truckType,
		PlateType:      plateType,
		ProductionYear: productionYear,
		Status:         status,
	}
	// validation input
	err := v.Struct(truckStruct)
	if err != nil {
		return res, err
	}
	DB.Where("id = ?", id).Updates(truckStruct)

	res.Status = http.StatusOK
	res.Message = "new truck successfully updated"
	res.Data = truckStruct

	return res, nil
}

func DeleteTruck(id int) (Response, error) {
	var res Response

	var truck Truck
	if err := DB.Where("id = ?", id).First(&truck).Error; err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "the truck which will be deleted not found !"
		return res, err
	}

	DB.Delete(&Truck{}, id)

	res.Status = http.StatusOK
	res.Message = "the truck successfully deleted"
	return res, nil
}
