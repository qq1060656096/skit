package db

import (
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

var ErrManagerNameNotFound = errors.New("db: db manager name not found")
var ErrManagerNameExist = errors.New("db: db manager name exist")
var ErrManagerNameOpen = errors.New("db: db manager name open error")

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

func (m *DbManager) Add(dbInfo *DbInfo) {
	if m.Exist(dbInfo.name){
		return
	}
	m.mutex.Lock()
	if !m.Exist(dbInfo.name) {
		m.data[dbInfo.name] = dbInfo
	}
	m.mutex.Unlock()
}

func (m *DbManager) Remove(name string) (*DbInfo, error) {
	if !m.Exist(name){
		return nil, ErrManagerNameNotFound
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	dbInfo, err := m.Get(name)
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
	return nil, ErrManagerNameNotFound
}

func (m *DbManager) Open(name, driverName, dataSourceName string) error {
	if m.Exist(name){
		return ErrManagerNameExist
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.Exist(name) {
		return ErrManagerNameExist
	}
	DbInfo, err := Open(name, driverName, dataSourceName)
	m.data[name] = DbInfo
	return err
}

func NewDbManager() *DbManager {
	return &DbManager{}
}

