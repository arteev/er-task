package routes

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"
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
