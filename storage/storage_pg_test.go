package storage

import (
	"fmt"
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
		Model: model.CarModel{
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

func TestFindCar(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer clearData(spg, t)
	_, err := s.FindCarByID(1)
	if err == nil {
		t.Error("Expected error")
	}
	car := addCar(spg, 1, t)
	got, err := s.FindCarByID(car.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != car.ID {
		t.Errorf("Expected car %v,got %v", car.ID, got.ID)
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
	err := s.Track("XXX", 0, 0)
	if err == nil {
		t.Error("Expected error")
	}
	want := `Car XXX not found`
	if got := err.Error(); got != want {
		t.Errorf("Expected error:%q, got %q", want, got)
	}

	//Add car with id=1
	car = addCar(spg, 1, t)
	if err := s.Track(car.Regnum, 1, 1); err != nil {
		t.Error(err)
	}
	if err := s.Track(car.Regnum, 1, 1); err != nil {
		t.Error(err)
	}
}
