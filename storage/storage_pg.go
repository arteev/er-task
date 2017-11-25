package storage

import (
	"database/sql"
	"fmt"

	"github.com/arteev/er-task/model"

	_ "github.com/lib/pq" //
)

var (
	sqlTrack  = `INSERT INTO "public"."LOCATION" ("CAR","POINT") VALUES($1,POINT($2,$3))`
	sqlAddCar = `INSERT INTO "CAR"("ID","MODEL","REGNUM") VALUES ($1,$2,$3)`
)

type storagePG struct {
	db         *sql.DB
	smttrac    *sql.Stmt
	stmtAddCar *sql.Stmt
}

func (pg *storagePG) Init(connection string) error {
	var err error
	pg.db, err = sql.Open("postgres", connection)
	if err != nil {
		return err
	}

	return pg.prepare()
}
func (pg *storagePG) Done() error {
	if pg.db == nil {
		return nil
	}
	return pg.db.Close()
}

func (pg *storagePG) prepare() (err error) {
	pg.smttrac, err = pg.db.Prepare(sqlTrack)
	if err != nil {
		return
	}
	pg.stmtAddCar, err = pg.db.Prepare(sqlAddCar)
	if err != nil {
		return
	}
	return nil
}

func (pg *storagePG) Track(car model.Car, latitude float64, longitude float64) error {
	_, err := pg.smttrac.Exec(&car.ID, &latitude, &latitude)
	return err
}

//TODO: find car by id
func (pg *storagePG) FindCarByID(id int) (*model.Car, error) {
	return nil, fmt.Errorf("Car %v not found", id)
}

func (pg *storagePG) addcar(car model.Car) error {
	_, err := pg.stmtAddCar.Exec(car.ID, car.Model.ID, car.Regnum)
	return err
}

func (pg *storagePG) addmodel(m model.ModelCar) error {

	return nil
}
