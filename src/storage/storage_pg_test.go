package storage

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/arteev/er-task/src/model"
)

//TODO: брать из env
var pgConnectionTest = "postgres://postgres:example@127.0.0.1/carrental?sslmode=disable"

func init() {
	if c, ok := os.LookupEnv("POSTGRES"); ok {
		pgConnectionTest = c
	}
}

func setUp(t *testing.T) Storage {
	t.Helper()
	s := GetStorage()
	err := s.Init(pgConnectionTest, false)
	if err != nil {
		t.Fatal(err)
	}
	spg := s.(*storagePG)
	sqls := []string{
		`INSERT INTO "CARTYPE"("ID","NAME","CODE") VALUES (1,'AUTO','AUTO')`,
		`INSERT INTO  "MODEL"("ID","NAME","CTYPE") VALUES (1,'test',1);`,
		`INSERT INTO  "DEPARTMENT"("ID","NAME") VALUES (1,'dep1');`,
	}
	for _, sql := range sqls {
		_, err = spg.db.Exec(sql)
		if err != nil {
			t.Errorf("Exec sql: %q. %s", sql, err)
		}
	}
	return s
}

func tearDown(s *storagePG, t *testing.T) {
	sqls := []string{
		`UPDATE "RENTAL" SET "DUMMY"=1`,
		`DELETE FROM "CARGOODS"`,
		`DELETE FROM "CARRENT"`,
		`DELETE FROM "RENTAL"`,
		`DELETE FROM "LOCATION"`,
		`DELETE FROM "CAR"`,
		`DELETE FROM "MODEL"`,
		`DELETE FROM "CARTYPE"`,
		`DELETE FROM "CAR"`,
		`DELETE FROM "DEPARTMENT"`,
	}
	for _, sql := range sqls {
		_, err := s.db.Exec(sql)
		if err != nil {
			t.Errorf("Exec sql: %q. %s", sql, err)
		}
	}
}

//for tests
func addcarstorage(pg *storagePG, car model.Car) error {
	_, err := pg.db.Exec(`INSERT INTO "CAR"("ID","MODEL","REGNUM") VALUES ($1,$2,$3)`, car.ID, car.Model.ID, car.Regnum)
	if err != nil {
		return err
	}
	_, err = pg.db.Exec(`INSERT INTO "CARGOODS"("DEPT","CAR") VALUES(1,$1)`, car.ID)
	return err
}

func addCar(s *storagePG, id int, t *testing.T) model.Car {
	t.Helper()
	c := model.Car{
		ID: id,
		Model: model.CarModel{
			ID:   1,
			Name: "test",
			//TODO: CARTYPE

		},
		Regnum: fmt.Sprintf("X%dXX", id),
	}
	if err := addcarstorage(s, c); err != nil {
		t.Fatal(err)

	}
	return c
}

func ExistsRentReturn(t *testing.T, s *storagePG, rn, dep, agent string, isRent bool) bool {
	op := 1
	if !isRent {
		op = 0
	}
	t.Helper()
	r := s.db.QueryRow(`SELECT COUNT(*) FROM "RENTAL" R,"CAR" C,"DEPARTMENT" d
	WHERE r."CAR" = c."ID"
	  and d."ID" = r."DEPT"	  
	  and c."REGNUM"= $1
	  and d."NAME" = $2
	  and r."AGENTNAME" = $3
	  and r."OPER"= $4
		 `, rn, dep, agent, op)
	var count int64
	if err := r.Scan(&count); err != nil {
		t.Error(err)
	}
	if count <= 0 {
		return false
	}
	//
	return true
}

////////////

func TestInitDonePg(t *testing.T) {
	s := GetStorage()
	if err := s.Done(); err != nil {
		t.Error(err)
	}
	err := s.Init("postgres://user:user@127.0.0.1/fake", false)
	if err == nil {
		t.Error("Expected error")
	}

	if err = s.Init(pgConnectionTest, false); err != nil {
		t.Fatal(err)
	}
	tearDown(s.(*storagePG), t)

	err = s.Done()
	if err != nil {
		t.Error(err)
	}
}

func TestFindCar(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer tearDown(spg, t)
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
	defer tearDown(spg, t)
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

	//TODO: Проверить данные добавлены или нет

}

func TestRentPg(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer tearDown(spg, t)

	//see:setUp
	_, err := s.Rent("000", "dep1", "000-000-000 01")
	if err == nil {
		t.Error("Expected error")
	}
	car := addCar(spg, 1, t)
	_, err = s.Rent(car.Regnum, "dep1", "000-000-000 01")
	if err != nil {
		t.Error(err)
	}
	if !ExistsRentReturn(t, spg, car.Regnum, "dep1", "000-000-000 01", true) {
		t.Error("Data not found for rent")
	}

	_, err = s.Return(car.Regnum, "dep1", "000-000-000 01")
	if err != nil {
		t.Error(err)
	}
	if !ExistsRentReturn(t, spg, car.Regnum, "dep1", "000-000-000 01", false) {
		t.Error("Data not found for return")
	}

	//test GetRentJornal
	rd, err := s.GetRentJornal("")
	if err != nil {
		t.Fatal(err)
	}

	if len(rd) != 2 {
		t.Fatalf("Expected count rent jornal %d, got %d", 2, len(rd))
	}

	want := "X1XX"
	if rd[0].RN != want {
		t.Errorf("Expected %q, got %q", want, rd[0].RN)
	}
	if rd[1].RN != want {
		t.Errorf("Expected %q, got %q", want, rd[1].RN)
	}
	if rd[0].Oper != "return" {
		t.Errorf("Expected %q, got %q", "return", rd[0].Oper)
	}
	if rd[1].Oper != "rent" {
		t.Errorf("Expected %q, got %q", "return", rd[1].Oper)
	}

	//GetRentJornal by car RN
	car2 := addCar(spg, 2, t)
	_, err = s.Rent(car2.Regnum, "dep1", "000-000-000 01")
	if err != nil {
		t.Error(err)
	}
	rd, err = s.GetRentJornal("X2XX")
	if err != nil {
		t.Fatal(err)
	}
	want = "X2XX"
	if len(rd) != 1 {
		t.Fatalf("Expected count rent jornal %d, got %d", 1, len(rd))
	}
	if rd[0].RN != want {
		t.Errorf("Expected %q, got %q", want, rd[0].RN)
	}

	//GetRentJornal not found
	rd, err = s.GetRentJornal("X3XX")
	if err != nil {
		t.Error(err)
	}
	if len(rd) != 0 {
		t.Fatalf("Expected count rent jornal %d, got %d", 0, len(rd))
	}
}

func TestCars(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer tearDown(spg, t)
	cars, err := s.GetCars()
	if err != nil {
		t.Fatal(err)
	}
	if len(cars) != 0 {
		t.Error("Expected empty list of cars")
	}
	addCar(spg, 1, t)
	addCar(spg, 2, t)
	cars, err = s.GetCars()
	if err != nil {
		t.Fatal(err)
	}
	if len(cars) == 0 {
		t.Errorf("Expected len(cars)=%d, got %d", 2, len(cars))
	}
	if want := "X1XX"; cars[0].Regnum != want {
		t.Errorf("Expected %q, got %q", cars[0].Regnum, want)
	}
	if want := "X2XX"; cars[1].Regnum != want {
		t.Errorf("Expected %q, got %q", cars[0].Regnum, want)
	}
	//TODO: check others fields
}

func TestDepartments(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer tearDown(spg, t)
	deps, err := s.GetDepartments()
	if err != nil {
		t.Fatal(err)
	}
	if len(deps) != 1 {
		t.Error("Expected len(list) = 1 of departments")
	}
	if deps[0].Name != "dep1" {
		t.Errorf("Expected name department %q, got %q", "dep1", deps[0].Name)
	}
}

func TestGetCarInfo(t *testing.T) {
	s := setUp(t)
	spg := s.(*storagePG)
	defer s.Done()
	defer tearDown(spg, t)
	_, err := s.GetCarInfo("00")
	if err != sql.ErrNoRows {
		t.Errorf("Expected %v, got %v", sql.ErrNoRows, err)
	}
	//see setUp()
	addCar(spg, 1, t)
	car, err := s.GetCarInfo(fmt.Sprintf("X%dXX", 1))
	if err != nil {
		t.Fatal(err)
	}
	if car.ID != 1 {
		t.Errorf("Expected %d,got %d", 1, car.ID)
	}
	if car.Department != "dep1" {
		t.Errorf("Expected %q,got %q", "dep1", car.Department)
	}
	if car.IsRental != 0 {
		t.Errorf("Expected %d,got %d", 0, car.IsRental)
	}
	//rent
	agnname := "000-000-000 01"
	_, err = s.Rent(car.Regnum, "dep1", agnname)
	if err != nil {
		t.Error(err)
	}
	car, err = s.GetCarInfo(fmt.Sprintf("X%dXX", 1))
	if err != nil {
		t.Fatal(err)
	}
	if car.IsRental != 1 {
		t.Errorf("Expected %d,got %d", 1, car.IsRental)
	}
	if car.Department != "dep1" {
		t.Errorf("Expected %q,got %q", "dep1", car.Department)
	}
	if car.Agent != agnname {
		t.Errorf("Expected %q,got %q", agnname, car.Agent)
	}
}
