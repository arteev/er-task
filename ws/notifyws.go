package ws

import "time"
import "github.com/arteev/er-task/storage"

type notifyRentWS struct {
	Type     string    `json:"type"`
	Model    string    `json:"model"`
	RN       string    `json:"rn"`
	Daterent time.Time `json:"daterent"`
	Dateret  time.Time `json:"dateret"`
	Agent    string    `json:"agent"`
	Oper     string    `json:"oper"`
}

func notifyRentFromStorage(n storage.Notification) *notifyRentWS {
	ret := &notifyRentWS{}
	if val, ok := n.Data["Type"]; ok {
		ret.Type = val.(string)
	}
	return ret
}
