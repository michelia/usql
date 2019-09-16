package usql

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var ErrNoRows = sql.ErrNoRows

var (
	Select = sq.Select
	Case   = sq.Case
	Expr   = sq.Expr
)

type (
	H   = map[string]interface{}
	Or  = sq.Or
	And = sq.And
	Eq  = sq.Eq
	Gl  = sq.Gt
)

// DB *sqlx.DB的简单封装
type DB struct {
	*sqlx.DB
}

// Insert squirrel 与 sqlx.db 结合
func (db *DB) Insert(into string, colVals H) (sql.Result, error) {
	return sq.Insert(into).SetMap(colVals).RunWith(db).Exec()
}

// Replace squirrel 与 sqlx.db 结合
func (db *DB) Replace(into string, colVals H) (sql.Result, error) {
	query, args, err := sq.Insert(into).SetMap(colVals).ToSql()
	if err != nil {
		return nil, err
	}
	query = "Replace" + query[6:]
	return db.Exec(query, args...)
}

// Delete squirrel 与 sqlx.db 结合
func (db *DB) Delete(into string) sq.DeleteBuilder {
	return sq.Delete(into).RunWith(db)
}

// Update squirrel 与 sqlx.db 结合
func (db *DB) Update(into string) sq.UpdateBuilder {
	return sq.Update(into).RunWith(db)
}

// SqGet squirrel 与 sqlx.Get 结合
func (db *DB) SqGet(columns ...string) sq.SelectBuilder {
	return sq.Select(columns...).RunWith(db)
}

// func (db *DB) SqGet(dest interface{}, selectBuilder sq.SelectBuilder) error {
// 	query, args, err := selectBuilder.ToSql()
// 	if err != nil {
// 		return err
// 	}
// 	return db.Get(dest, query, args...)
// }

// SqSelect squirrel 与 sqlx.Select 结合
func (db *DB) SqSelect(dest interface{}, selectBuilder sq.SelectBuilder) error {
	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return err
	}
	return db.Select(dest, query, args...)
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
// https://jmoiron.github.io/sqlx/#namedParams
// p := Place{TelephoneCode: 50}
// pp := []Place{}
// // select all telcodes > 50
// nstmt, err := db.PrepareNamed(`SELECT * FROM place WHERE telcode > :telcode`)
// err = nstmt.Select(&pp, p)
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
