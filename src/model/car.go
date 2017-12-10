package model

import "time"

type Car struct {
	ID     int      `json:"-"`
	Regnum string   `json:"rn"`
	Model  CarModel `json:"model"`
}

type CarInfo struct {
	Car
	Department string    `json:"department"`
	Dateoper   time.Time `json:"dateoper"`
	IsRental   int       `json:"isrent"`
	Agent      string    `json:"agent"`
}
