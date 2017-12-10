package model

import (
	"time"
)

type (
	Car struct {
		ID     int      `json:"-"`
		Regnum string   `json:"rn"`
		Model  CarModel `json:"model"`
	}

	CarInfo struct {
		Car
		Department string    `json:"department"`
		Dateoper   time.Time `json:"dateoper"`
		IsRental   int       `json:"isrent"`
		Agent      string    `json:"agent"`
	}

	//CarModel - Модель ТС
	CarModel struct {
		ID   int     `json:"-"`
		Name string  `json:"name"`
		Type CarType `json:"cartype"`
	}

	//CarType тип ТС
	CarType struct {
		ID   int    `json:"-"`
		Name string `json:"type"`
		Code string `json:"code"`
	}

	//Department - подразделения проката ТС
	Department struct {
		ID   int    `json:"-"`
		Name string `json:"name"`
	}

	StatsItem struct {
		Entity      string        `json:"entity"`
		Count       int           `json:"count"`
		Duration    time.Duration `json:"-"`
		DurationStr string        `json:"avgduration"`
	}

	StatsDepartment struct {
		Department string      `json:"department"`
		Stats      []StatsItem `json:"entities"`
	}

	RentData struct {
		Type     string    `json:"type"`
		Model    string    `json:"model"`
		RN       string    `json:"rn"`
		Dateoper time.Time `json:"dateoper"`
		Agent    string    `json:"agent"`
		Oper     string    `json:"oper"`
		Dept     string    `json:"dept"`
	}

	Response struct {
		ContentType string `json:"content,omitempty"`
		Error       string `json:"error,omitempty"`
		Message     string `json:"message,omitempty"`
	}
	RentDataResponse struct {
		Response
		Data []RentData `json:"data"`
	}

	CarsResponse struct {
		Response
		Data []Car `json:"data"`
	}
	CarInfoResponse struct {
		Response
		Data CarInfo `json:"data"`
	}

	DepartmentsResponse struct {
		Response
		Data []Department `json:"data"`
	}

	StatsDepartmentoResponse struct {
		Response
		Data []StatsDepartment `json:"data"`
	}
)
