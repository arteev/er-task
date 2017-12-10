package tests

import (
	"errors"
	"fmt"
	"sync"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
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
	invokedCars        bool
	invokedCarInfo     bool
	invokedDepartments bool

	sync.RWMutex
	cars           map[string]model.Car
	carid          map[int]model.Car
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

func (s *FakeStorage) Rent(rn string, dep string, agn string) (int, error) {
	s.invokedRent = true
	s.Lock()
	defer s.Unlock()
	car, exist := s.cars[rn]
	if !exist {
		return 0, errors.New("Car not found")
	}
	d, exist := s.department[dep]
	if !exist {
		return 0, errors.New("Department not found")
	}

	rj := model.RentData{
		RN:    rn,
		Agent: agn,
		//TODO: Dep
		//Dep: dep
	}
	s.rentjournal[car.Regnum+d.Name+agn+"Rent"] = rj
	s.rentjournalArr = append(s.rentjournalArr, rj)

	return len(s.rentjournalArr), nil
}

func (s *FakeStorage) Return(rn string, dep string, agn string) (int, error) {
	s.invokedReturn = true
	s.Lock()
	defer s.Unlock()
	car, exist := s.cars[rn]
	if !exist {
		return 0, errors.New("Car not found")
	}
	d, exist := s.department[dep]
	if !exist {
		return 0, errors.New("Department not found")
	}
	rj := model.RentData{
		RN:    rn,
		Agent: agn,
		//TODO: Dep
		//Dep: dep
	}
	s.rentjournal[car.Regnum+d.Name+agn+"Return"] = rj
	s.rentjournalArr = append(s.rentjournalArr, rj)
	return len(s.rentjournalArr), nil
}

func (s *FakeStorage) GetRentJornal(rn string) ([]model.RentData, error) {
	s.Lock()
	defer s.Unlock()
	s.invokedRentJournal = true
	rds := make([]model.RentData, 0)

	for i := len(s.rentjournalArr) - 1; i >= 0; i-- {
		if s.rentjournalArr[i].RN == rn || rn == "" {
			rds = append(rds, s.rentjournalArr[i])
		}

	}
	return rds, nil
}

func (s *FakeStorage) Notify() chan storage.Notification {
	//TODO:
	return nil
}

func (s *FakeStorage) GetCars() ([]model.Car, error) {
	s.Lock()
	defer s.Unlock()
	s.invokedCars = true
	//WARN: выдает записи не в той последовательности
	cars := make([]model.Car, 0)
	for _, car := range s.cars {
		cars = append(cars, car)
	}
	return cars, nil
}

func (s *FakeStorage) GetCarInfo(rn string) (*model.CarInfo, error) {
	s.Lock()
	defer s.Unlock()
	s.invokedCarInfo = true
	car, exists := s.cars[rn]
	if !exists {
		return nil, fmt.Errorf("Car %q not found", rn)
	}
	ci := &model.CarInfo{
		Car: car,
	}
	return ci, nil
}

//Статистика в разрезе подразделений и моделей
func (s *FakeStorage) GetStatsByModel() ([]model.StatsDepartment, error) {
	return nil, nil
}

//Статистика в разрезе подразделений и тип ТС
func (s *FakeStorage) GetStatsByType() ([]model.StatsDepartment, error) {
	return nil, nil

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
	//TODO: type car
	s.carmodel[name] = model.CarModel{ID: id, Name: name}
	return s.carmodel[name]
}

//helper for test. Add/Update Car
func (s *FakeStorage) addcar(id int, rn string, m model.CarModel) model.Car {
	s.Lock()
	defer s.Unlock()
	s.cars[rn] = model.Car{ID: id, Regnum: rn, Model: m}
	s.carid[id] = model.Car{ID: id, Regnum: rn, Model: m}
	return s.cars[rn]
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
	s.carmodel = make(map[string]model.CarModel)
	s.rentjournal = make(map[string]model.RentData)
	s.rentjournalArr = make([]model.RentData, 0)
	s.track = make(map[string][]point)
}

func (s *FakeStorage) GetDepartments() ([]model.Department, error) {
	s.RLock()
	defer s.RUnlock()
	s.invokedDepartments = true
	deps := make([]model.Department, 0)
	for _, d := range s.department {
		deps = append(deps, d)
	}
	return deps, nil
}
