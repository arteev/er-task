package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arteev/er-task/storage"
)

func assertCodeEqual(t *testing.T, text string, want, got int) {
	t.Helper()
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
	if len(b) == 0 {
		return &response{}, nil
	}
	val := response{}
	err = json.Unmarshal(b, &val)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func checkResponseJSON(t *testing.T, r io.Reader, got, want string, near bool) {
	t.Helper()
	if !near && got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
	if near && !strings.Contains(got, want) {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func checkResponseJSONError(t *testing.T, r io.Reader, want string, near bool) {
	t.Helper()
	val, err := json2response(r)
	if err != nil {
		t.Error(err)
		return
	}
	checkResponseJSON(t, r, val.Error, want, near)
}

func checkResponseJSONMessage(t *testing.T, r io.Reader, want string, near bool) {
	t.Helper()
	val, err := json2response(r)
	if err != nil {
		t.Error(err)
		return
	}
	checkResponseJSON(t, r, val.Message, want, near)
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
	//Invoked
	if !fakestorage.invokedTrack {
		t.Error("Must be invoke Storage.Track")
	}

	assertCodeEqual(t, "Expected:Car not found", http.StatusNotFound, w.Code)
	checkResponseJSONError(t, w.Body, `Car 0 not found`, false)

	//tracking the car
	r, _ = http.NewRequest("PUT", "/api/v1/tracking/1/55.755/37.6251", nil)
	w = httptest.NewRecorder()
	a.routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Expected:Success", http.StatusOK, w.Code)
	checkResponseJSONMessage(t, w.Body, `success`, false)

	//TODO: проверить какие данные вставлены при трекинге
}
