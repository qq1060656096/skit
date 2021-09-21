package db

import (
	"testing"
)

func TestOpen(t *testing.T) {
	dbInfo, err := Open("testOpen", SQLite, "/Users/zhaoweijie/go/src/github.com/qq1060656096/skit/test/data/db.testOpen.sqlite3.db")
	if err != nil {
		t.Error("open sqlite 3 fail.", err)
	}
	sql := `
DROP TABLE "sqlite3_open2" ;
CREATE TABLE "sqlite3_open2" (
  "name" TEXT
);
`
	db := dbInfo.DB()
	defer dbInfo.Close()
	_, err = db.Exec(sql)
	if err != nil {
		t.Error("open sqlite3 create database fail.", err)
	}
}