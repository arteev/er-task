package model

import "time"

type RentData struct {
	Type     string    `json:"type"`
	Model    string    `json:"model"`
	RN       string    `json:"rn"`
	Dateoper time.Time `json:"dateoper"`
	Agent    string    `json:"agent"`
	SS       string    `json:"ss"`
	Oper     string    `json:"oper"`
}
