package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/arteev/er-task/storage"
)

func TestRentAPI(t *testing.T) {
	//TODO:Invoked test
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = &FakeStorage{}
		return fakestorage
	}
	a := new(App)
	a.init()
	//http.Handle("/", a.init())
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

	//Car not found
}
