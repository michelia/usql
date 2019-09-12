package usql

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrNoRows = sql.ErrNoRows

// DB *sqlx.DB的简单封装
type DB struct {
	*sqlx.DB
}

// NamedGet 类似于 Get
func (db *DB) NamedGet(dest interface{}, query string, arg interface{}) error {
	nstmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = nstmt.Get(dest, arg)
	return err
}

// NamedSelect 类似于 Select
func (db *DB) NamedSelect(dest interface{}, query string, arg interface{}) error {
	nstmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = nstmt.Select(dest, arg)
	return err
}

// Setting
// Lifetime 连接的生命期 单位是分钟
// idle 最大闲置数
func (db *DB) Setting(Lifetime, idle int) {
	// https://colobu.com/2019/05/27/configuring-sql-DB-for-better-performance/
	db.SetConnMaxLifetime(time.Minute * time.Duration(Lifetime)) // 处理 Driver: invalid connection
	db.SetMaxIdleConns(idle)
	// db.SetMaxOpenConns(n)
}

func Connect(driverName, dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	d := &DB{DB: db}
	d.Setting(10, 1)
	return d, nil
}

func MustConnect(driverName, dataSourceName string) *DB {
	db := sqlx.MustConnect(driverName, dataSourceName)
	d := &DB{DB: db}
	d.Setting(10, 1)
	return d
}
