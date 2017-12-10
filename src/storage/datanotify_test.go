package storage

import "testing"

func TestRentFromStorage(t *testing.T) {
	c := Notification{
		Name: "",
		Data: map[string]interface{}{
			"DATEOPER": "2017-12-02T09:50:54.696989",
			"TYPE":     "AUTO",
		},
	}
	got := RentDataFromStorage(c)
	if got.Dateoper.Format(pglayout) != c.Data["DATEOPER"] {
		t.Errorf("Expected %q, got %q", c.Data["DATEOPER"], got.Dateoper.Format(pglayout))
	}
	if got.Type != c.Data["TYPE"] {
		t.Errorf("Eepected %q, got %q", c.Data["TYPE"], got.Type)
	}
	//etc...
}
