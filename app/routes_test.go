package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/arteev/er-task/model"
	"github.com/arteev/er-task/storage"
)

/*
type Storage interface {
	Init(string) error
	Done() error

	//Трекинг ТС по координатам GPS. Возможно нужна высота?
	Track(model.Car, float64, float64) error

	FindCarByID(id int) (model.Car, error)
}
*/
type FakeStorage struct {
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
		Regnum: "XX001X",
	}
	return nil
}

func (s *FakeStorage) Done() error {
	return nil
}

func (s *FakeStorage) Track(model.Car, float64, float64) error {
	return nil
}

func (s *FakeStorage) FindCarByID(id int) (*model.Car, error) {
	s.RLock()
	defer s.RUnlock()
	car, exists := s.cars[id]
	if !exists {
		return nil, fmt.Errorf("Car %v not found", id)
	}
	return &car, nil
}

func initFakeStorage() storage.Storage {
	return &FakeStorage{}
}

//////////////////////////////

func assertCodeEqual(t *testing.T, text string, want, got int) {
	if want != got {
		t.Errorf("%s: Expected http.code %d, got %d", text, want, got)
	}
}

func TestTrackAPI(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = &FakeStorage{}
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()

	//invalid path(Variables empty)
	r, _ := http.NewRequest("PUT", "/api/v1/tracking/0/0", nil)
	w := httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "invalid path", http.StatusNotFound, w.Code)

	//car not found
	r, _ = http.NewRequest("PUT", "/api/v1/tracking/0/55.755/37.6251", nil)
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Expected:Car not found", http.StatusNotFound, w.Code)

	//tracking the car
	r, _ = http.NewRequest("PUT", "/api/v1/tracking/1/55.755/37.6251", nil)
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Expected:Success", http.StatusOK, w.Code)

	got := w.Body.String()
	want := `{"message":"success"}`
	if got != want {
		t.Errorf("Expected %q,got %q", want, got)
	}
}
