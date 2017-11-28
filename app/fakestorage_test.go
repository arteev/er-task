package app

import (
	"errors"
	"fmt"
	"sync"

	"github.com/arteev/er-task/model"
	"github.com/arteev/er-task/storage"
)

type FakeStorage struct {
	invokedTrack    bool
	invokedFindByID bool
	invokedRent     bool
	invokedReturn   bool

	sync.RWMutex
	cars map[int]model.Car
}

func (s *FakeStorage) Init(string) error {
	s.Lock()
	defer s.Unlock()
	s.cars = make(map[int]model.Car)
	s.cars[1] = model.Car{
		ID:     1,
		Model:  model.ModelCar{ID: 1, Name: "test"},
		Regnum: "XX1X",
	}
	return nil
}

func (s *FakeStorage) Done() error {
	return nil
}

func (s *FakeStorage) Track(rn string, x float64, y float64) error {
	s.invokedTrack = true
	if rn == "0" {
		return errors.New("Car 0 not found")
	}
	return nil
}

func (s *FakeStorage) FindCarByID(id int) (*model.Car, error) {
	s.RLock()
	defer s.RUnlock()
	s.invokedFindByID = true
	car, exists := s.cars[id]
	if !exists {
		return nil, fmt.Errorf("Car %v not found", id)
	}
	return &car, nil
}

//TODO:!
func (s *FakeStorage) Rent(rn string, dep string, agn string) error {
	s.invokedRent = true
	return errors.New("Car not found")
}

//TODO:!
func (s *FakeStorage) Return(rn string, dep string, agn string) error {
	s.invokedReturn = true
	return errors.New("Car not found")
}

func initFakeStorage() storage.Storage {
	return &FakeStorage{}
}
