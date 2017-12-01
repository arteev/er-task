package ws

import "time"
import "github.com/arteev/er-task/storage"

type notifyRentWS struct {
	Type     string    `json:"type"`
	Model    string    `json:"model"`
	RN       string    `json:"rn"`
	Dateoper time.Time `json:"dateoper"`
	Agent    string    `json:"agent"`
	SS       string    `json:"ss"`
	Oper     string    `json:"oper"`
}

func notifyRentFromStorage(n storage.Notification) *notifyRentWS {
	ret := &notifyRentWS{}
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
	if val, ok := n.Data["AGENTCODE"]; ok {
		ret.SS = val.(string)
	}
	if val, ok := n.Data["OPER"]; ok {
		ret.Oper = val.(string)
	}
	if val, ok := n.Data["DATEOPER"]; ok {
		var err error
		ret.Dateoper, err = time.Parse(time.RFC3339, val.(string))
		if err != nil {
			//log it
		}
	}
	return ret
}
