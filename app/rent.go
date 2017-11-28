package app

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"
)

type rentdata struct {
	RegNum   string `scheme:"regnum"` // Номер ТС
	DeptCode string `scheme:"dept"`   // Код подразделения
	AgentSS  string `scheme:"agent"`  // Код(ИД) агента
}

func (r rentdata) Check() error {
	if r.RegNum == "" || r.DeptCode == "" || r.AgentSS == "" {
		return errors.New("Value not found")
	}
	return nil
}

func newRentData(r *http.Request) (*rentdata, int, error) {
	if err := r.ParseForm(); err != nil {
		return nil, http.StatusBadRequest, err
	}
	rd := new(rentdata)
	decoder := schema.NewDecoder()
	err := decoder.Decode(rd, r.Form)
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
