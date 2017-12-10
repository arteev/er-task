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
	ContentType string `json:"content,omitempty"`
	Error       string `json:"error,omitempty"`
	Message     string `json:"message,omitempty"`
}
type RentDataResponse struct {
	Response
	Data []RentData `json:"data"`
}

type CarsResponse struct {
	Response
	Data []Car `json:"data"`
}
type CarInfoResponse struct {
	Response
	Data CarInfo `json:"data"`
}

type DepartmentsResponse struct {
	Response
	Data []Department `json:"data"`
}

type StatsItem struct {
	Entity      string        `json:"entity"`
	Count       int           `json:"count"`
	Duration    time.Duration `json:"-"`
	DurationStr string        `json:"avgduration"`
}
type StatsDepartment struct {
	Department string      `json:"department"`
	Stats      []StatsItem `json:"entities"`
}

type StatsDepartmentoResponse struct {
	Response
	Data []StatsDepartment `json:"data"`
}
