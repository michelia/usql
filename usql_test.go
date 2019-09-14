package usql

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
)

func TestSql(t *testing.T) {
	a, b, c := Select("c1, c2, c3").Columns("a").From("t use index(key)").Where("c1 = 3").Where(Or{Expr("c2=4"), Eq{"c4": 5}}).ToSql()
	fmt.Println(a, ":", b, c)
	a, b, c = sq.Insert("t").Columns("a,b,c").Values(1, 2, 3).ToSql()
	a = "Replace" + a[6:]
	fmt.Println(a, ":", b, c)
}
