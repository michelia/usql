package usql

import (
	"testing"

	"github.com/gocraft/dbr/v2"
)

func TestSql(t *testing.T) {
	t.Log(SqlStr(
		dbr.Select("column, `a`").Where("query=3 and a in ?", []int64{1, 2, 3, 4, 5}).Where("aa=?", 9).From("tab use index(xxx)").From("new use index()"),
	))
	t.Log(SqlStr(
		dbr.InsertInto("a").Columns("c1", "c2").Values(1, 2).Values(6, "abc'`"),
	))
	t.Log(SqlStr(
		dbr.Select("count(id)").From(
			dbr.Select("*").From("suggestions").As("t_count"),
		),
	))
	t.Log(SqlStr(
		dbr.Select("*").From("suggestions").
			Join("subdomains", "suggestions.subdomain_id = subdomains.id").
			Join("accounts", "subdomains.accounts_id = accounts.id").Where(dbr.Or(dbr.Eq("abc", 3), dbr.Neq("b", "4"))).Where("time > abc"),
	))
	t.Log(SqlStr(
		dbr.Update("user").Set("a", 3).Set("b", dbr.Expr("? + ?", dbr.I("b"), 3)).Set("time", dbr.Expr("now()")),
	))
}
