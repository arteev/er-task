package storage

import (
	"github.com/arteev/er-task/model"
)

type Storage interface {
	Init(string) error
	Done() error

	//Трекинг ТС с рег.номером по координатам GPS. Возможно нужна высота?
	Track(rn string, x float64, y float64) error

	//Взять в аренду ТС
	Rent(rn string, dep string, agn string) error
	//Вернуть ТС
	Return(rn string, dep string, agn string) error

	//Поиск ТС по ID
	FindCarByID(id int) (*model.Car, error)
}

//Для тестирования переопределить
var GetStorage = getStorageDefault

func getStorageDefault() Storage {
	return &storagePG{}
}
