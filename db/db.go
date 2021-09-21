package db

import (
	"database/sql"
)


type DbInfo struct {
	name           string
	driverName     string
	dataSourceName string
	db             *sql.DB
}

func (d *DbInfo) DB() *sql.DB {
	return d.db
}

func (d *DbInfo) Close() error {
	return d.db.Close()
}

func Open(name, driverName, dataSourceName string) (*DbInfo, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	dbInfo := &DbInfo{
		name:       name,
		driverName: driverName,
		db:         db,
	}
	return dbInfo, nil
}

func OpenDB(name, driverName string, db *sql.DB) *DbInfo {
	dbInfo := &DbInfo{
		name: name,
		driverName: driverName,
		db: db,
	}
	return dbInfo
}


