package storage

import (
	"fmt"
	"strings"
	"testing"

	"github.com/arteev/er-task/model"
)

//TODO: брать из env
var pgConnectionTest = "postgres://postgres:example@192.168.1.43/carrental?sslmode=disable"

func setUp(t *testing.T) Storage {
	s := GetStorage()
	err := s.Init(pgConnectionTest)
	if err != nil {
		t.Fatal(err)
	}
	spg := s.(*storagePG)
	_, err = spg.db.Exec(`INSERT INTO "CARTYPE"("ID","NAME","CODE") VALUES (1,'AUTO','AUTO')`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = spg.db.Exec(`INSERT INTO  "MODEL"("ID","NAME","CTYPE") VALUES (1,'test',1);`)
	if err != nil {
		t.Fatal(err)
	}

	return s
}

func clearData(s *storagePG, t *testing.T) {

	_, err := s.db.Exec(`DELETE FROM "LOCATION"`)
	if err != nil {
		t.Error(err)
	}

	_, err = s.db.Exec(`DELETE FROM "CAR"`)
	if err != nil {
		t.Error(err)
	}

	_, err = s.db.Exec(`DELETE FROM "MODEL"`)
	if err != nil {
		t.Error(err)
	}

	_, err = s.db.Exec(`DELETE FROM "CARTYPE"`)
	if err != nil {
		t.Error(err)
	}

	_, err = s.db.Exec(`DELETE FROM "CAR"`)
	if err != nil {
		t.Error(err)
	}

}

func addCar(s *storagePG, id int, t *testing.T) model.Car {
	//model = test
	c := model.Car{
		ID: id,
		Model: model.ModelCar{
			ID:   1,
			Name: "test",
			//TODO: CARTYPE

		},
		Regnum: fmt.Sprintf("X%dXX", id),
	}
	if err := s.addcar(c); err != nil {
		t.Fatal(err)

	}
	return c
}

////////////

func TestInitDonePg(t *testing.T) {
	s := GetStorage()
	if err := s.Done(); err != nil {
		t.Error(err)
	}
	err := s.Init("postgres://user:user@127.0.0.1/fake")
	if err == nil {
		t.Error("Expected error")
	}

	if err = s.Init(pgConnectionTest); err != nil {
		t.Fatal(err)
	}
	clearData(s.(*storagePG), t)

	err = s.Done()
	if err != nil {
		t.Error(err)
	}

}
func TestTrackPg(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer clearData(spg, t)
	//empty DB
	car := model.Car{}

	//test unknown car
	err := s.Track(car, 0, 0)
	if err == nil {
		t.Error("Expected error")
	}
	if got := err.Error(); !strings.Contains(strings.ToUpper(got), "LOCATION_CAR_FKEY") {
		want := `insert or update on table "LOCATION" violates foreign key constraint "LOCATION_CAR_fkey"`
		t.Errorf("Expected error:%q, got %q", want, got)
	}

	//Add car with id=1
	car = addCar(spg, 1, t)
	if err := s.Track(car, 1, 1); err != nil {
		t.Error(err)
	}
	if err := s.Track(car, 1, 1); err != nil {
		t.Error(err)
	}
}
