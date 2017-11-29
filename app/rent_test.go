package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/arteev/er-task/storage"
)

func TestRentAndReturnAPI(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()

	for _, routeTest := range []struct {
		route   string
		name    string
		invoked *bool
	}{
		{"/api/v1/rent", "Rent", &fakestorage.invokedRent},
		{"/api/v1/return", "Return", &fakestorage.invokedReturn},
	} {
		fakestorage.clear()
		//form values empty
		r, _ := http.NewRequest("POST", routeTest.route, nil)
		w := httptest.NewRecorder()
		a.routes.ServeHTTP(w, r)
		assertCodeEqual(t, "Form value not missing", http.StatusBadRequest, w.Code)
		checkResponseJSONError(t, w.Body, `missing form body`, false)

		//Value not found
		form := url.Values{}
		r, _ = http.NewRequest("POST", routeTest.route, strings.NewReader(form.Encode()))
		w = httptest.NewRecorder()
		a.routes.ServeHTTP(w, r)
		assertCodeEqual(t, "Form value not found", http.StatusBadRequest, w.Code)
		checkResponseJSONError(t, w.Body, `Value not found`, true)

		//Car no found
		form = url.Values{}
		form.Set("regnum", "000")
		form.Set("dept", "000")
		form.Set("agent", "000")
		r, _ = http.NewRequest("POST", routeTest.route, strings.NewReader(form.Encode()))
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
		w = httptest.NewRecorder()
		a.routes.ServeHTTP(w, r)
		//Invoked
		if !*routeTest.invoked {
			t.Errorf("Must be invoke Storage.%s", routeTest.name)
		}
		assertCodeEqual(t, "Car not found", http.StatusInternalServerError, w.Code)
		checkResponseJSONError(t, w.Body, `Car not found`, true)

		//Rent/Return
		dep := fakestorage.adddepart(1, "dep1")
		md := fakestorage.addmodel(1, "bmw")
		car := fakestorage.addcar(1, "X000XX", md)
		agn := fakestorage.addagent(1, "000-000-000 01", "иван", "иванович", "иванов")
		form = url.Values{}
		form.Set("regnum", car.Regnum)
		form.Set("dept", dep.Name)
		form.Set("agent", agn.Code)
		r, _ = http.NewRequest("POST", routeTest.route, strings.NewReader(form.Encode()))
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
		w = httptest.NewRecorder()
		a.routes.ServeHTTP(w, r)
		assertCodeEqual(t, routeTest.name, http.StatusOK, w.Code)
		checkResponseJSONMessage(t, w.Body, `success`, false)
		if rentok := fakestorage.existsRent(car.Regnum, dep.Name, agn.Code, routeTest.name); !rentok {
			t.Errorf("Expected %s in storage", routeTest.name)
		}
	}
}
