package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arteev/er-task/src/storage"
)

func TestTrackAPI(t *testing.T) {
	var fakestorage *FakeStorage
	storage.GetStorage = func() storage.Storage {
		fakestorage = initFakeStorage().(*FakeStorage)
		return fakestorage
	}
	routes := GetRoutes(storage.GetStorage())
	defer fakestorage.Done()

	//invalid path(Variables empty)
	r, _ := http.NewRequest("PUT", "/api/v1/tracking/0/0", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, r)
	assertCodeEqual(t, "invalid path", http.StatusNotFound, w.Code)

	//car not found
	r, _ = http.NewRequest("PUT", "/api/v1/tracking/0/55.755/37.6251", nil)
	w = httptest.NewRecorder()
	routes.ServeHTTP(w, r)
	//Invoked
	if !fakestorage.invokedTrack {
		t.Error("Must be invoke Storage.Track")
	}

	assertCodeEqual(t, "Expected:Car not found", http.StatusNotFound, w.Code)
	checkResponseJSONError(t, w.Body, `Car 0 not found`, false)

	//tracking the car
	r, _ = http.NewRequest("PUT", "/api/v1/tracking/1/55.755/37.6251", nil)
	w = httptest.NewRecorder()
	routes.ServeHTTP(w, r)
	assertCodeEqual(t, "Expected:Success", http.StatusOK, w.Code)
	checkResponseJSONMessage(t, w.Body, `success`, false)

	//TODO: проверить какие данные вставлены при трекинге
	if got := fakestorage.countTrack("1"); got != 1 {
		t.Errorf("Expected count track %d, got %d", 1, got)
	}
}
