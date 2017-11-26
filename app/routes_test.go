package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/arteev/er-task/model"
	"github.com/arteev/er-task/storage"
)

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
		Regnum: "XX1X",
	}
	return nil
}

func (s *FakeStorage) Done() error {
	return nil
}

func (s *FakeStorage) Track(rn string, x float64, y float64) error {
	if rn == "0" {
		return errors.New("Car 0 not found")
	}
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

//TODO: refactor this: helper package
type response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func json2response(r io.Reader) (*response, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	val := response{}
	err = json.Unmarshal(b, &val)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func checkResponseJSONError(t *testing.T, r io.Reader, want string) {
	t.Helper()
	val, err := json2response(r)
	if err != nil {
		t.Error(err)
		return
	}
	if val.Error != want {
		t.Errorf("Expected %q, got %q", want, val.Error)
	}
}

func checkResponseJSONMessage(t *testing.T, r io.Reader, want string) {
	t.Helper()
	val, err := json2response(r)
	if err != nil {
		t.Error(err)
		return
	}
	if val.Message != want {
		t.Errorf("Expected %q, got %q", want, val.Error)
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
	checkResponseJSONError(t, w.Body, `Car 0 not found`)

	//tracking the car
	r, _ = http.NewRequest("PUT", "/api/v1/tracking/1/55.755/37.6251", nil)
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Expected:Success", http.StatusOK, w.Code)
	checkResponseJSONMessage(t, w.Body, `success`)

	//todo: проверить какие данные вставлены при трекинге
}
