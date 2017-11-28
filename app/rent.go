package app

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"
)

type RentData struct {
	RegNum   string `schema:"regnum"` // Номер ТС
	DeptCode string `schema:"dept"`   // Код подразделения
	AgentSS  string `schema:"agent"`  // Код(ИД) агента
}

func (r RentData) Check() error {
	if r.RegNum == "" || r.DeptCode == "" || r.AgentSS == "" {
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

func (a *App) Rent(w http.ResponseWriter, r *http.Request) (int, error) {
	rd, code, err := newRentData(r)
	if err != nil {
		return code, err
	}
	err = a.db.Rent(rd.RegNum, rd.DeptCode, rd.AgentSS)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (a *App) Return(w http.ResponseWriter, r *http.Request) (int, error) {
	rd, code, err := newRentData(r)
	if err != nil {
		return code, err
	}
	err = a.db.Return(rd.RegNum, rd.DeptCode, rd.AgentSS)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
