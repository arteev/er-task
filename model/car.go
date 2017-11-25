package model

type ModelCar struct {
	ID   int
	Name string
}
type Car struct {
	ID     int
	Regnum string
	Model  ModelCar
}
