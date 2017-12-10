package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
)

func TestDepartments(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()

	fakestorage.adddepart(1, "dep1")
	r, _ := http.NewRequest("GET", "/api/v1/departments", nil)
	w := httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "", http.StatusOK, w.Code)
	//Invoked
	if !fakestorage.invokedDepartments {
		t.Error("Must be invoke Storage.Departments")
	}
	deps := &model.DepartmentsResponse{}
	err := json.Unmarshal(w.Body.Bytes(), deps)
	if err != nil {
		t.Fatal(err)
	}
	if deps.Message != "success" {
		t.Errorf("Expected message %q, got %q", "success", deps.Message)
	}
	if deps.ContentType != "departments" {
		t.Errorf("Expected ContentType %q, got %q", "departments", deps.ContentType)
	}
	if deps.Data == nil {
		t.Fatal("Expected Data not nil")
	}
	if len(deps.Data) == 0 {
		t.Fatal("Expected len([]deps) > 0")
	}
	if deps.Data[0].Name != "dep1" {
		t.Errorf("Expected %q, got %q", "dep1", deps.Data[0].Name)
	}
}
