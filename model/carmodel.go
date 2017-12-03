package model

//CarModel - Модель ТС
type CarModel struct {
	ID   int     `json:"-"`
	Name string  `json:"name"`
	Type CarType `json:"cartype"`
}

//CarType тип ТС
type CarType struct {
	ID   int    `json:"-"`
	Name string `json:"type"`
	Code string `json:"code"`
}
