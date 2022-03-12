package models

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Driver struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required" gorm:"not null"`
	Phone     string    `json:"phone" validate:"required" gorm:"not null"`
	IdCard    string    `json:"id_card" validate:""`
	License   string    `json:"license" validate:""`
	Status    bool      `json:"status" validate:"" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" validate:""`
}

func FetchAllDrivers() (Response, error) {
	var res Response

	drivers := []Driver{}

	DB.Find(&drivers)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = drivers

	return res, nil
}

func FetchByIdDriver(id int) (Response, error) {
	var res Response

	driver := Driver{}

	if err := DB.Where("id = ?", id).First(&driver).Error; err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "the driver not found !"
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = driver

	return res, nil
}

func StoreDriver(licenseNumber, phone, idCard, license string) (Response, error) {
	var res Response

	v := validator.New()
	driverStruct := Driver{
		Name:    licenseNumber,
		Phone:   phone,
		IdCard:  idCard,
		License: license,
		Status:  true,
	}
	// validation input
	err := v.Struct(driverStruct)
	if err != nil {
		return res, err
	}

	DB.Create(&driverStruct)

	res.Status = http.StatusOK
	res.Message = "new driver successfully created"
	res.Data = driverStruct
	return res, nil
}

func UpdateDriver(id int, licenseNumber, phone, idCard, license string, status bool) (Response, error) {
	var res Response

	var driver Driver
	if err := DB.Where("id = ?", id).First(&driver).Error; err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "the driver which will be updated not found !"
		return res, err
	}

	v := validator.New()
	driverStruct := Driver{
		Id:      id,
		Name:    licenseNumber,
		Phone:   phone,
		IdCard:  idCard,
		License: license,
		Status:  status,
	}
	// validation input
	err := v.Struct(driverStruct)
	if err != nil {
		return res, err
	}
	DB.Where("id = ?", id).Updates(driverStruct)

	res.Status = http.StatusOK
	res.Message = "new driver successfully updated"
	res.Data = driverStruct

	return res, nil
}

func DeleteDriver(id int) (Response, error) {
	var res Response

	var driver Driver
	if err := DB.Where("id = ?", id).First(&driver).Error; err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "the driver which will be deleted not found !"
		return res, err
	}

	DB.Delete(&Driver{}, id)

	res.Status = http.StatusOK
	res.Message = "the driver successfully deleted"
	return res, nil
}
