package storage

import (
	"github.com/arteev/er-task/model"
)

type Storage interface {
	Init(string) error
	Done() error

	//Трекинг ТС с рег.номером по координатам GPS. Возможно нужна высота?
	Track(string, float64, float64) error

	//Поиск ТС по ID
	FindCarByID(id int) (*model.Car, error)
}

//Для тестирования переопределить
var GetStorage = getstoragePostgres

func getstoragePostgres() Storage {
	return &storagePG{}
}
