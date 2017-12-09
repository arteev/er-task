package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lib/pq"

	"github.com/arteev/er-task/model"

	_ "github.com/lib/pq" //
)

const (
	opReturn = 0
	opRent   = 1
)

const (
	channel    = "webnotify"
	pingPeriod = 60 * time.Second
)

//TODO: refactor template Repository

type storagePG struct {
	connection string
	listener   *pq.Listener
	notify     chan Notification

	db *sql.DB

	smttrac          *sql.Stmt
	stmtFindCarByID  *sql.Stmt
	stmtRentAction   *sql.Stmt
	stmtRentJornal   *sql.Stmt
	stmtCars         *sql.Stmt
	stmtDepartments  *sql.Stmt
	stmtCarInfo      *sql.Stmt
	stmtStatsByModel *sql.Stmt
	stmtStatsByType  *sql.Stmt
}

func (pg *storagePG) Init(connection string, usenotify bool) error {
	var err error

	pg.connection = connection
	pg.db, err = sql.Open("postgres", connection)
	if err != nil {
		return err
	}
	if usenotify {
		pg.notify = make(chan Notification, 50)
		if err := pg.initlistener(); err != nil {
			close(pg.notify)
			return err
		}
	}
	return pg.prepare()
}
func (pg *storagePG) Done() error {
	if pg.db == nil {
		return nil
	}
	if pg.listener != nil {
		pg.listener.UnlistenAll()
		close(pg.notify)
	}
	return pg.db.Close()
}

func (pg *storagePG) initlistener() error {
	pgevent := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println("pgbroadcast: ", err.Error())
		}
	}
	pg.listener = pq.NewListener(pg.connection, 10*time.Second, time.Minute, pgevent)
	go pg.handleNotifications()
	if err := pg.listener.Listen(channel); err != nil {
		return err
	}
	return nil
}

func (pg *storagePG) handleNotifications() {
	//TODO: cancel
	for {
		select {
		case n := <-pg.listener.Notify:
			if n == nil {
				continue
			}
			var notify Notification
			err := json.Unmarshal([]byte(n.Extra), &notify)
			if err != nil {
				log.Println("Error JSON: ", err)
			} else {
				fmt.Println("notify: ", notify)
				go func() { pg.notify <- notify }()
			}
		case <-time.After(pingPeriod):
			go func() {
				pg.listener.Ping()
			}()
		}
	}
}

func (pg *storagePG) prepare() (err error) {
	pg.smttrac, err = pg.db.Prepare(sqlTrack)
	if err != nil {
		return err
	}
	pg.stmtFindCarByID, err = pg.db.Prepare(sqlFindByID)
	if err != nil {
		return err
	}
	pg.stmtRentAction, err = pg.db.Prepare(sqlRentAction)
	if err != nil {
		return err

	}
	pg.stmtRentJornal, err = pg.db.Prepare(sqlRentJornal)
	if err != nil {
		return err
	}
	pg.stmtCars, err = pg.db.Prepare(sqlCars)
	if err != nil {
		return err
	}
	pg.stmtDepartments, err = pg.db.Prepare(sqlDepartments)
	if err != nil {
		return err
	}
	pg.stmtCarInfo, err = pg.db.Prepare(sqlCarInfo)
	if err != nil {
		return err
	}

	pg.stmtStatsByModel, err = pg.db.Prepare(sqlStatsByModel)
	if err != nil {
		return err
	}

	pg.stmtStatsByType, err = pg.db.Prepare(sqlStatsByType)
	if err != nil {
		return err
	}
	return nil
}

//Трекинг ТС с рег.номером по координатам GPS.
func (pg *storagePG) Track(regnum string, latitude float64, longitude float64) error {
	//TODO: ??В очередь т.к. HL. если да то надо выше декоратор юзать??
	_, err := pg.smttrac.Exec(regnum, latitude, longitude)
	if err != nil && strings.Contains(err.Error(), `"CAR" violates not-null constraint`) {
		return fmt.Errorf("Car %s not found", regnum)
	}
	return err
}

//Взять в аренду ТС
func (pg *storagePG) Rent(rn string, dep string, agn string) (int, error) {
	var id int
	err := pg.stmtRentAction.QueryRow(opRent, rn, dep, agn).Scan(&id)
	if err != nil {
		return 0, err
	}
	//TODO: parse text error. Replace text error
	return id, nil
}

//Вернуть ТС
func (pg *storagePG) Return(rn string, dep string, agn string) (int, error) {
	var id int
	err := pg.stmtRentAction.QueryRow(opReturn, rn, dep, agn).Scan(&id)
	if err != nil {
		return 0, err
	}
	//TODO: parse text error. Replace text error
	return id, nil
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

func (pg *storagePG) Notify() chan Notification {
	return pg.notify
}

//TODO: test it 1 car
func (pg *storagePG) GetRentJornal(rn string) ([]model.RentData, error) {
	rows, err := pg.stmtRentJornal.Query(rn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rds := make([]model.RentData, 0)
	for rows.Next() {
		var id int
		r := &model.RentData{}
		err := rows.Scan(&id, &r.Type, &r.Model, &r.RN, &r.Dateoper, &r.Agent, &r.Oper, &r.Dept)
		if err != nil {
			return nil, err
		}
		rds = append(rds, *r)
	}
	return rds, nil
}

func (pg *storagePG) GetCars() ([]model.Car, error) {
	rows, err := pg.stmtCars.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cars := make([]model.Car, 0)
	for rows.Next() {
		car := model.Car{}
		//TODO : IDs
		err := rows.Scan(&car.ID, &car.Regnum, &car.Model.Name, &car.Model.Type.Code, &car.Model.Type.Name)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

//TODO : трансформировать ошибку
//TODO: test it
func (pg *storagePG) GetCarInfo(rn string) (*model.CarInfo, error) {
	row := pg.stmtCarInfo.QueryRow(rn)
	car := &model.CarInfo{}
	err := row.Scan(&car.ID, &car.Regnum, &car.Model.Name, &car.Model.Type.Name, &car.Model.Type.Code, &car.Department, &car.Dateoper, &car.IsRental, &car.Agent)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (pg *storagePG) GetDepartments() ([]model.Department, error) {
	rows, err := pg.stmtDepartments.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	deps := make([]model.Department, 0)
	for rows.Next() {
		d := model.Department{}
		err = rows.Scan(&d.ID, &d.Name)
		if err != nil {
			return nil, err
		}
		deps = append(deps, d)
	}
	return deps, nil
}

func (pg *storagePG) getstats(s *sql.Stmt) ([]model.StatsDepartment, error) {
	rows, err := s.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sts := make([]*model.StatsDepartment, 0)
	deps := make(map[string]*model.StatsDepartment, 0)
	for rows.Next() {
		var (
			depname, entity string
			count           int
			dur             float64
		)

		err := rows.Scan(&depname, &entity, &count, &dur)
		if err != nil {
			return nil, err
		}
		s, exists := deps[depname]
		if !exists {
			s = &model.StatsDepartment{
				Stats: make([]model.StatsItem, 0),
			}
			deps[depname] = s
			sts = append(sts, s)
		}
		s.Department = depname
		s.Stats = append(s.Stats, model.StatsItem{
			Count:       count,
			Duration:    time.Duration(dur * float64(time.Second)),
			DurationStr: time.Duration(dur * float64(time.Second) / float64(count)).String(),
			Entity:      entity,
		})

	}
	res := make([]model.StatsDepartment, len(sts))
	for i, s := range sts {
		res[i] = *s
	}
	return res, nil

}

//TODO: test it
//Статистика в разрезе подразделений и моделей
func (pg *storagePG) GetStatsByModel() ([]model.StatsDepartment, error) {
	return pg.getstats(pg.stmtStatsByModel)
}

//TODO: test it
//Статистика в разрезе подразделений и тип ТС
func (pg *storagePG) GetStatsByType() ([]model.StatsDepartment, error) {
	return pg.getstats(pg.stmtStatsByType)
}
