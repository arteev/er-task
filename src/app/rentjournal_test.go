package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
)

func TestRentJournal(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()

	dep := fakestorage.adddepart(1, "dep1")
	md := fakestorage.addmodel(1, "bmw")
	car := fakestorage.addcar(1, "X000XX", md)
	agn := "000-000-000 01"
	_, err := fakestorage.Rent(car.Regnum, dep.Name, agn)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := http.NewRequest("GET", "/api/v1/rentjournal", nil)
	w := httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusOK, w.Code)
	//Invoked
	if !fakestorage.invokedRentJournal {
		t.Error("Must be invoke Storage.RentJournal")
	}

	rds := &model.RentDataResponse{}
	err = json.Unmarshal(w.Body.Bytes(), rds)
	if err != nil {
		t.Fatal(err)
	}
	if rds.Message != "success" {
		t.Errorf("Expected message %q, got %q", "success", rds.Message)
	}
	if rds.Data == nil {
		t.Error("Expected not nil []RentData")
	}
	if len(rds.Data) == 0 {
		t.Fatal("Expected len([]RentData) > 0")
	}
	if rds.Data[0].RN != car.Regnum {
		t.Errorf("Expected %q,got %q", car.Regnum, rds.Data[0].RN)
	}
	//second car
	car2 := fakestorage.addcar(2, "X002XX", md)
	_, err = fakestorage.Rent(car2.Regnum, dep.Name, agn)
	if err != nil {
		t.Fatal(err)
	}
	r, _ = http.NewRequest("GET", "/api/v1/rentjournal", nil)
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusOK, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), rds)
	if err != nil {
		t.Fatal(err)
	}
	if got := len(rds.Data); got != 2 {
		t.Errorf("Expected len([]RentData) == 2, got %d", got)
	}
}

func TestRentJournalByRNCar(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()
	dep := fakestorage.adddepart(1, "dep1")
	md := fakestorage.addmodel(1, "bmw")
	car := fakestorage.addcar(1, "X000XX", md)
	agn := "000-000-000 01"
	_, err := fakestorage.Rent(car.Regnum, dep.Name, agn)
	if err != nil {
		t.Fatal(err)
	}
	//test not found
	r, _ := http.NewRequest("GET", "/api/v1/rentjournal/000", nil)
	w := httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusNotFound, w.Code)

	r, _ = http.NewRequest("GET", "/api/v1/rentjournal/X000XX", nil)
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusOK, w.Code)
	rds := &model.RentDataResponse{}
	err = json.Unmarshal(w.Body.Bytes(), rds)
	if err != nil {
		t.Fatal(err)
	}
	if rds.Message != "success" {
		t.Errorf("Expected message %q, got %q", "success", rds.Message)
	}
	if len(rds.Data) == 0 {
		t.Fatal("Expected len([]RentData) > 0")
	}
	if rds.Data[0].RN != car.Regnum {
		t.Errorf("Expected %q,got %q", car.Regnum, rds.Data[0].RN)
	}
}
