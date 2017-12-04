package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arteev/er-task/model"
	"github.com/arteev/er-task/storage"
)

func TestCars(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()
	md := fakestorage.addmodel(1, "bmw")
	car := fakestorage.addcar(1, "XXX", md)
	r, _ := http.NewRequest("GET", "/api/v1/cars", nil)
	w := httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
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
