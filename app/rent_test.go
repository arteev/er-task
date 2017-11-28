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

func TestRentAPI(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = &FakeStorage{}
		return fakestorage
	}
	a := new(App)
	a.init()
	defer fakestorage.Done()

	//form values empty
	r, _ := http.NewRequest("POST", "/api/v1/rent", nil)
	w := httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Form value not missing", http.StatusBadRequest, w.Code)
	checkResponseJSONError(t, w.Body, `missing form body`, false)

	//Value not found
	form := url.Values{}
	r, _ = http.NewRequest("POST", "/api/v1/rent", strings.NewReader(form.Encode()))
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Form value not found", http.StatusBadRequest, w.Code)
	checkResponseJSONError(t, w.Body, `Value not found`, true)

	//Car no found
	form = url.Values{}
	form.Set("regnum", "000")
	form.Set("dept", "000")
	form.Set("agent", "000")
	r, _ = http.NewRequest("POST", "/api/v1/rent", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)

	//Invoked
	if !fakestorage.invokedRent {
		t.Error("Must be invoke Storage.Rent")
	}

	assertCodeEqual(t, "Car not found", http.StatusInternalServerError, w.Code)
	checkResponseJSONError(t, w.Body, `Car not found`, true)

}
