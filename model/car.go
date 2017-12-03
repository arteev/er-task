package model

type Car struct {
	ID     int      `json:"-"`
	Regnum string   `json:"rn"`
	Model  CarModel `json:"model"`
}
