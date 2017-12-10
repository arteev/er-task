package routes

import (
	"errors"
	"net/http"

	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/schema"
)

type RentData struct {
	RegNum   string `schema:"regnum"` // Номер ТС
	DeptCode string `schema:"dept"`   // Код подразделения
	Agent    string `schema:"agent"`  // Фио агента
}

func (r RentData) Check() error {
	if r.RegNum == "" || r.DeptCode == "" || r.Agent == "" {
		return errors.New("Value not found")
	}
	return nil
}

func newRentData(r *http.Request) (*RentData, int, error) {
	if err := r.ParseForm(); err != nil {
		return nil, http.StatusBadRequest, err
	}
	rd := new(RentData)
	decoder := schema.NewDecoder()
	err := decoder.Decode(rd, r.PostForm)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	err = rd.Check()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return rd, 0, nil
}

func Rent(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	rd, code, err := newRentData(r)
	if err != nil {
		return nil, code, err
	}
	_, err = db.Rent(rd.RegNum, rd.DeptCode, rd.Agent)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &struct {
		Message string `json:"message"`
	}{
		"success",
	}, http.StatusOK, nil
}

func Return(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	rd, code, err := newRentData(r)
	if err != nil {
		return nil, code, err
	}
	_, err = db.Return(rd.RegNum, rd.DeptCode, rd.Agent)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &struct {
		Message string `json:"message"`
	}{
		"success",
	}, http.StatusOK, nil
}
