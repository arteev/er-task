package storage

import "testing"

func TestRentFromStorage(t *testing.T) {
	c := Notification{
		Name: "",
		Data: map[string]interface{}{
			"DATEOPER": "2017-12-02T09:50:54.696989",
			"TYPE":     "AUTO",
			"MODEL":    "VAZ",
			"AGENT":    "JAMES BOND",
			"OPER":     "rent",
			"RN":       "001",
			"DEPT":     "dep1",
		},
	}
	got, err := RentDataFromStorage(c)
	if err != nil {
		t.Fatal(err)
	}
	if got.Dateoper.Format(pglayout) != c.Data["DATEOPER"] {
		t.Errorf("Expected %q, got %q", c.Data["DATEOPER"], got.Dateoper.Format(pglayout))
	}
	if got.Type != c.Data["TYPE"] {
		t.Errorf("Eepected %q, got %q", c.Data["TYPE"], got.Type)
	}
	if got.Model != c.Data["MODEL"] {
		t.Errorf("Eepected %q, got %q", c.Data["MODEL"], got.Model)
	}
	if got.Agent != c.Data["AGENT"] {
		t.Errorf("Eepected %q, got %q", c.Data["AGENT"], got.Agent)
	}
	if got.Oper != c.Data["OPER"] {
		t.Errorf("Eepected %q, got %q", c.Data["OPER"], got.Oper)
	}
	if got.Dept != c.Data["DEPT"] {
		t.Errorf("Eepected %q, got %q", c.Data["DEPT"], got.Dept)
	}

	c = Notification{
		Data: map[string]interface{}{
			"DATEOPER": "1232139",
		},
	}
	got, err = RentDataFromStorage(c)
	if err == nil {
		t.Errorf("Expected error, got error nil")
	}
}
