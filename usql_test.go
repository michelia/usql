package usql

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/gocraft/dbr/v2"
)

func TestSql(t *testing.T) {
	// a, b, c := Select("c1, c2, c3").Columns("a").From("t use index(key)").Where("c1 = 3").Where(Or{Expr("c2=4"), Eq{"c4": 5}}).ToSql()
	// fmt.Println(a, ":", b, c)
	a, b, c := sq.Insert("t").Columns("a,b,c").Values(1, 2, 3).SetMap(map[string]interface{}{
		"a": 4,
		"b": 5,
		"c": 6,
	}).ToSql()
	// a = "Replace" + a[6:]
	// a, b, c = sq.Select("a1, a2").Where(`a=3`).From("ttt").ToSql()
	fmt.Println(a, ":", b, c)
	s := dbr.Select("column, `a`").Where("query=3").Where("aa=9").From("tab use index(xxx)").From("new use index()")
	t.Logf("xxx: %+v\n", SqlStr(s))
	i := dbr.InsertInto("a").Pair("xxx", 123)
	t.Logf("xxx: %+v\n", SqlStr(i))
}
