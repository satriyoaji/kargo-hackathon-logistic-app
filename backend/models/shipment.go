package models

type Shipment struct {
	ShipmentNumber string           `json:"shipment_number"`
	LicenseNumber  string           `json:"license_number"`
	DriverName     string           `json:"driver_name"`
	Origin         string           `json:"origin"`
	Destination    string           `json:"destination"`
	LoadingDate    string           `json:"loading_date"`
	Status         []ShipmentStatus `json:"status"`
	Action         []ShipmentAction `json:"action"`
}

type ShipmentStatus struct {
	Status string `json:"status"`
}

type ShipmentAction struct {
	Action string `json:"action"`
}
