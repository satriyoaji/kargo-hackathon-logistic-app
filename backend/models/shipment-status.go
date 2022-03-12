package models

import (
	"net/http"
)

type ShipmentStatus struct {
	Status string `json:"status"`
}

func GetAllStatus() (Response, error) {
	var res Response

	shipmentStatus := []ShipmentStatus{}
	shipmentStatus = append(shipmentStatus,
		ShipmentStatus{"Ongoing to Origin"},
		ShipmentStatus{"At Origin"},
		ShipmentStatus{"Ongoing to Destination"},
		ShipmentStatus{"At Destination"},
		ShipmentStatus{"Completed"},
	)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = shipmentStatus

	return res, nil
}
