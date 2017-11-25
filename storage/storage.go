package storage

import (
	"github.com/arteev/er-task/model"
)

type Storage interface {
	Init(string) error
	Done() error

	//Трекинг ТС по координатам GPS. Возможно нужна высота?
	Track(model.Car, float64, float64) error

	FindCarByID(id int) (*model.Car, error)
}

//Для тестирования переопределить
var GetStorage = getstoragePostgres

func getstoragePostgres() Storage {
	return &storagePG{}
}
