package storage

import (
	"time"

	"github.com/arteev/er-task/src/model"
)

const pglayout = "2006-01-02T15:04:05.999999"

//RentDataFromStorage returns model.RentData from notification storage
func RentDataFromStorage(n Notification) (*model.RentData, error) {
	ret := &model.RentData{}
	if val, ok := n.Data["TYPE"]; ok {
		ret.Type = val.(string)
	}
	if val, ok := n.Data["MODEL"]; ok {
		ret.Model = val.(string)
	}
	if val, ok := n.Data["RN"]; ok {
		ret.RN = val.(string)
	}
	if val, ok := n.Data["AGENT"]; ok {
		ret.Agent = val.(string)
	}
	if val, ok := n.Data["OPER"]; ok {
		ret.Oper = val.(string)
	}
	if val, ok := n.Data["DATEOPER"]; ok {
		var err error
		ret.Dateoper, err = time.Parse(pglayout, val.(string))
		if err != nil {
			return nil, err
		}
	}
	if val, ok := n.Data["DEPT"]; ok {
		ret.Dept = val.(string)
	}
	return ret, nil
}
