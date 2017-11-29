package storage

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/arteev/er-task/model"

	_ "github.com/lib/pq" //
)

const (
	opReturn = 0
	opRent   = 1
)

var (
	sqlTrack = `INSERT INTO "public"."LOCATION" ("CAR","POINT") VALUES(  
		(SELECT "ID" FROM "CAR" WHERE "REGNUM"=$1),
		POINT($2,$3))`
	sqlAddCar   = `INSERT INTO "CAR"("ID","MODEL","REGNUM") VALUES ($1,$2,$3)`
	sqlFindByID = `SELECT c."ID", c."REGNUM", c."MODEL", m."NAME" "MODELNAME" FROM "CAR" c ,"MODEL" m
	WHERE
		c."MODEL" = m."ID"
		and c."ID"=$1`

	sqlRentJornal = `INSERT INTO "RENTAL"("TSWORK","OPER","CAR","DEPT","AGENT") VALUES (
		CURRENT_TIMESTAMP,$1,
		(SELECT "CAR"."ID" from "CAR" WHERE  "CAR"."REGNUM" = $2),
		(SELECT "DEPARTMENT"."ID" FROM "DEPARTMENT" WHERE "DEPARTMENT"."NAME" = $3),
		(SELECT "AGENT"."ID" FROM "AGENT" WHERE "AGENT"."CODE" = $4))`
)

//TODO: refactor template Repository

type storagePG struct {
	db              *sql.DB
	smttrac         *sql.Stmt
	stmtAddCar      *sql.Stmt
	stmtFindCarByID *sql.Stmt
	stmtRentJornal  *sql.Stmt
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
		return err
	}
	pg.stmtAddCar, err = pg.db.Prepare(sqlAddCar)
	if err != nil {
		return err
	}

	pg.stmtFindCarByID, err = pg.db.Prepare(sqlFindByID)
	if err != nil {
		return err
	}
	pg.stmtRentJornal, err = pg.db.Prepare(sqlRentJornal)
	if err != nil {
		return err

	}
	return nil
}

//Трекинг ТС с рег.номером по координатам GPS.
func (pg *storagePG) Track(regnum string, latitude float64, longitude float64) error {
	//TODO: ??В очередь т.к. HL. если да то надо выше декоратор юзать??
	_, err := pg.smttrac.Exec(regnum, latitude, longitude)
	if err != nil && strings.Contains(err.Error(), `"CAR" violates not-null constrain`) {
		return fmt.Errorf("Car %s not found", regnum)
	}
	return err
}

//Взять в аренду ТС
//TODO:!
func (pg *storagePG) Rent(rn string, dep string, agn string) error {
	_, err := pg.stmtRentJornal.Exec(opRent, rn, dep, agn)
	if err != nil {
		return err
	}
	return nil
}

//Вернуть ТС
//TODO:!
func (pg *storagePG) Return(rn string, dep string, agn string) error {
	_, err := pg.stmtRentJornal.Exec(opReturn, rn, dep, agn)
	if err != nil {
		return err
	}
	return nil
}

//TODO: refactor this. REPO
func (pg *storagePG) FindCarByID(id int) (*model.Car, error) {
	row := pg.stmtFindCarByID.QueryRow(id)
	var (
		idc            sql.NullInt64
		idmodel        sql.NullInt64
		regnum, smodel string
	)
	err := row.Scan(&idc, &regnum, &idmodel, &smodel)
	if err != nil {
		return nil, err
	}
	if !idc.Valid {
		return nil, fmt.Errorf("Car %v not found", id)
	}
	car := &model.Car{
		ID:     int(idc.Int64),
		Regnum: regnum,
		Model: model.CarModel{
			ID:   int(idmodel.Int64),
			Name: smodel,
		},
	}
	return car, err
}

func (pg *storagePG) addcar(car model.Car) error {
	_, err := pg.stmtAddCar.Exec(car.ID, car.Model.ID, car.Regnum)
	return err
}

func (pg *storagePG) addmodel(m model.CarModel) error {
	return nil
}
