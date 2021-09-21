package db

import (
	"database/sql"
	"errors"
	"sync"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
	SQLServer = "mssql"
)

var DefaultDbManger = NewDbManager()

var errDbManagerNameNotFound = errors.New("db: db manager name not found")
var errDbManagerNameExist = errors.New("db: db manager name exist")
var errDbManagerNameOpen = errors.New("db: db manager name open error")

type DbManager struct {
	mutex sync.Mutex
	data map[string]*DbInfo
}

func (m *DbManager) Exist(name string) bool {
	if _,ok := m.data[name]; ok {
		return true
	}
	return false
}

func (m *DbManager) Add(dbInfo *DbInfo) error {
	if m.Exist(dbInfo.name){
		return errDbManagerNameExist
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.Exist(dbInfo.name){
		return errDbManagerNameExist
	}
	m.data[dbInfo.name] = dbInfo
	return nil
}

func (m *DbManager) Remove(name string) (*DbInfo, error) {
	dbInfo, err := m.Get(name)
	if err != nil {
		return nil, err
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	dbInfo, err = m.Get(name)
	if err != nil {
		return nil, err
	}
	delete(m.data, name)
	return dbInfo, nil
}


func (m *DbManager) Get(name string) (*DbInfo, error) {
	if dbInfo, ok := m.data[name]; ok {
		return dbInfo, nil
	}
	return nil, errDbManagerNameNotFound
}

func (m *DbManager) Open(name, driverName, dataSourceName string) error {
	if m.Exist(name){
		return errDbManagerNameExist
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.Exist(name) {
		return errDbManagerNameExist
	}
	DbInfo, err := Open(name, driverName, dataSourceName)
	m.data[name] = DbInfo
	return err
}

func (m *DbManager) OpenDB(name, driverName string, db *sql.DB) error {
	if m.Exist(name){
		return errDbManagerNameExist
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.Exist(name) {
		return errDbManagerNameExist
	}
	DbInfo := OpenDB(name, driverName, db)
	m.data[name] = DbInfo
	return nil
}

func NewDbManager() *DbManager {
	return &DbManager{}
}

