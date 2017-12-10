package model

//Department - подразделения проката ТС
type Department struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
}
