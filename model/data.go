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

type RentDataResponse struct {
	Error   string     `json:"error"`
	Message string     `json:"message"`
	Data    []RentData `json:"data"`
}
