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
	cars        map[string]model.Car
	carid       map[int]model.Car
	agent       map[string]model.Agent
	department  map[string]model.Department
	carmodel    map[string]model.CarModel
	rentjournal map[string]struct{}
}

func (s *FakeStorage) Init(string) error {
	s.Lock()
	defer s.Unlock()
	/*s.cars[1] = model.Car{
		ID:     1,
		Model:  model.ModelCar{ID: 1, Name: "test"},
		Regnum: "XX1X",
	}*/
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
	car, exists := s.carid[id]
	if !exists {
		return nil, fmt.Errorf("Car %v not found", id)
	}
	return &car, nil
}

func (s *FakeStorage) existsRent(rn string, dep string, agn string, suffix string) bool {
	_, exist := s.rentjournal[rn+dep+agn+suffix]
	return exist
}

func (s *FakeStorage) Rent(rn string, dep string, agn string) error {
	s.invokedRent = true
	s.Lock()
	defer s.Unlock()
	car, exist := s.cars[rn]
	if !exist {
		return errors.New("Car not found")
	}
	d, exist := s.department[dep]
	if !exist {
		return errors.New("Department not found")
	}
	a, exist := s.agent[agn]
	if !exist {
		return errors.New("Agent not found")
	}
	s.rentjournal[car.Regnum+d.Name+a.Code+"Rent"] = struct{}{}
	return nil
}

func (s *FakeStorage) Return(rn string, dep string, agn string) error {
	s.invokedReturn = true
	s.Lock()
	defer s.Unlock()
	car, exist := s.cars[rn]
	if !exist {
		return errors.New("Car not found")
	}
	d, exist := s.department[dep]
	if !exist {
		return errors.New("Department not found")
	}
	a, exist := s.agent[agn]
	if !exist {
		return errors.New("Agent not found")
	}
	s.rentjournal[car.Regnum+d.Name+a.Code+"Return"] = struct{}{}
	return nil
}

func initFakeStorage() storage.Storage {
	return &FakeStorage{
		cars:        make(map[string]model.Car),
		carid:       make(map[int]model.Car),
		department:  make(map[string]model.Department),
		agent:       make(map[string]model.Agent),
		carmodel:    make(map[string]model.CarModel),
		rentjournal: make(map[string]struct{}),
	}
}

//helper for test. Add/Update Department
func (s *FakeStorage) adddepart(id int, name string) model.Department {
	s.Lock()
	defer s.Unlock()
	s.department[name] = model.Department{id, name}
	return s.department[name]
}

//helper for test. Add/Update Model
func (s *FakeStorage) addmodel(id int, name string) model.CarModel {
	s.Lock()
	defer s.Unlock()
	s.carmodel[name] = model.CarModel{id, name}
	return s.carmodel[name]
}

//helper for test. Add/Update Car
func (s *FakeStorage) addcar(id int, rn string, m model.CarModel) model.Car {
	s.Lock()
	defer s.Unlock()
	s.cars[rn] = model.Car{id, rn, m}
	s.carid[id] = model.Car{id, rn, m}
	return s.cars[rn]
}

//helper for test. Add/Update Agent
func (s *FakeStorage) addagent(id int, code, name, midname, family string) model.Agent {
	s.Lock()
	defer s.Unlock()
	s.agent[code] = model.Agent{id, code, name, family, midname}
	return s.agent[code]
}

func (s *FakeStorage) clear() {
	s.cars = make(map[string]model.Car)
	s.carid = make(map[int]model.Car)
	s.department = make(map[string]model.Department)
	s.agent = make(map[string]model.Agent)
	s.carmodel = make(map[string]model.CarModel)
	s.rentjournal = make(map[string]struct{})
}
