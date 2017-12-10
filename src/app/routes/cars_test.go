package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
)

func TestCarInfo(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	routes := GetRoutes(storage.GetStorage())
	defer fakestorage.Done()

	md := fakestorage.addmodel(1, "bmw")
	fakestorage.addcar(1, "XXX", md)

	//Invoked
	r, _ := http.NewRequest("GET", "/api/v1/car/0", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusInternalServerError, w.Code)
	if !fakestorage.invokedCarInfo {
		t.Error("Must be invoke Storage.CarInfo")
	}

	r, _ = http.NewRequest("GET", "/api/v1/car/XXX", nil)
	w = httptest.NewRecorder()
	routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusOK, w.Code)
	ci := &model.CarInfoResponse{}
	err := json.Unmarshal(w.Body.Bytes(), ci)
	if err != nil {
		t.Fatal(err)
	}
	if ci.Message != "success" {
		t.Errorf("Expected message %q, got %q", "success", ci.Message)
	}
	if ci.Data.Car.Regnum != "XXX" {
		t.Errorf("Expected %q,got %q", "XXX", ci.Data.Car.Regnum)
	}
}

func TestCars(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	routes := GetRoutes(storage.GetStorage())

	defer fakestorage.Done()
	md := fakestorage.addmodel(1, "bmw")
	car := fakestorage.addcar(1, "XXX", md)
	r, _ := http.NewRequest("GET", "/api/v1/cars", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusOK, w.Code)
	//Invoked
	if !fakestorage.invokedCars {
		t.Error("Must be invoke Storage.Cars")
	}
	cars := &model.CarsResponse{}
	err := json.Unmarshal(w.Body.Bytes(), cars)
	if err != nil {
		t.Fatal(err)
	}
	if cars.Message != "success" {
		t.Errorf("Expected message %q, got %q", "success", cars.Message)
	}
	if cars.ContentType != "cars" {
		t.Errorf("Expected ContentType %q, got %q", "cars", cars.ContentType)
	}
	if cars.Data == nil {
		t.Fatal("Expected Data not nil")
	}
	if len(cars.Data) == 0 {
		t.Fatal("Expected len([]Cars) > 0")
	}
	if cars.Data[0].Regnum != car.Regnum {
		t.Errorf("Expected %q, got %q", car.Regnum, cars.Data[0].Regnum)
	}
}
