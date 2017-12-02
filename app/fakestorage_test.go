package app

import (
	"errors"
	"fmt"
	"sync"

	"github.com/arteev/er-task/model"
	"github.com/arteev/er-task/storage"
)

type point struct {
	latitude  float64
	longitude float64
}

type FakeStorage struct {
	invokedTrack       bool
	invokedFindByID    bool
	invokedRent        bool
	invokedReturn      bool
	invokedRentJournal bool

	sync.RWMutex
	cars           map[string]model.Car
	carid          map[int]model.Car
	agent          map[string]model.Agent
	department     map[string]model.Department
	carmodel       map[string]model.CarModel
	rentjournal    map[string]model.RentData
	rentjournalArr []model.RentData
	track          map[string][]point
}

func initFakeStorage() storage.Storage {
	return &FakeStorage{
		cars:           make(map[string]model.Car),
		carid:          make(map[int]model.Car),
		department:     make(map[string]model.Department),
		agent:          make(map[string]model.Agent),
		carmodel:       make(map[string]model.CarModel),
		rentjournal:    make(map[string]model.RentData),
		rentjournalArr: make([]model.RentData, 0),
		track:          make(map[string][]point),
	}
}

func (s *FakeStorage) Init(string, bool) error {
	s.Lock()
	defer s.Unlock()
	return nil
}

func (s *FakeStorage) Done() error {
	return nil
}

func (s *FakeStorage) Track(rn string, latitude float64, longitude float64) error {
	s.Lock()
	defer s.Unlock()
	s.invokedTrack = true
	if rn == "0" {
		return errors.New("Car 0 not found")
	}
	c, exists := s.track[rn]
	if !exists {
		c = make([]point, 0)
	}
	c = append(c, point{latitude, longitude})
	s.track[rn] = c
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

	rj := model.RentData{
		RN:    rn,
		Agent: agn,
		//TODO: Dep
		//Dep: dep
	}
	s.rentjournal[car.Regnum+d.Name+a.Code+"Rent"] = rj
	s.rentjournalArr = append(s.rentjournalArr, rj)

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
	rj := model.RentData{
		RN:    rn,
		Agent: agn,
		//TODO: Dep
		//Dep: dep
	}
	s.rentjournal[car.Regnum+d.Name+a.Code+"Return"] = rj
	s.rentjournalArr = append(s.rentjournalArr, rj)
	return nil
}

func (s *FakeStorage) GetRentJornal() ([]model.RentData, error) {
	s.Lock()
	defer s.Unlock()
	s.invokedRentJournal = true
	rds := make([]model.RentData, 0)
	for i := len(s.rentjournalArr) - 1; i >= 0; i-- {
		rds = append(rds, s.rentjournalArr[i])
	}
	return rds, nil
}

func (pg *FakeStorage) Notify() chan storage.Notification {
	//TODO:
	return nil
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

//helper for test. Count track coordinates by regnum of the car
func (s *FakeStorage) countTrack(rn string) int {
	s.RLock()
	defer s.RUnlock()
	c, exists := s.track[rn]
	if !exists {
		return 0
	}
	return len(c)
}

func (s *FakeStorage) clear() {
	s.cars = make(map[string]model.Car)
	s.carid = make(map[int]model.Car)
	s.department = make(map[string]model.Department)
	s.agent = make(map[string]model.Agent)
	s.carmodel = make(map[string]model.CarModel)
	s.rentjournal = make(map[string]model.RentData)
	s.rentjournalArr = make([]model.RentData, 0)
	s.track = make(map[string][]point)
}
