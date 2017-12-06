package model

import (
	"time"
)

type RentData struct {
	Type     string    `json:"type"`
	Model    string    `json:"model"`
	RN       string    `json:"rn"`
	Dateoper time.Time `json:"dateoper"`
	Agent    string    `json:"agent"`
	SS       string    `json:"ss"`
	Oper     string    `json:"oper"`
	Dept     string    `json:"dept"`
}

type Response struct {
	ContentType string `json:"content"`
	Error       string `json:"error"`
	Message     string `json:"message"`
}
type RentDataResponse struct {
	Response
	Data []RentData `json:"data"`
}

type CarsResponse struct {
	Response
	Data []Car `json:"data"`
}

type DepartmentsResponse struct {
	Response
	Data []Department `json:"data"`
}
