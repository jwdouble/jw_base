package sqlc

import (
	"cmd/main.go/sqlc/gen"
	"jw.lib/sqlx"
)

var Q *gen.Queries

func init() {
	sqlx.Register(sqlx.DefaultSqlDriver, sqlx.DefaultSqlAddr)
	db := sqlx.GetSqlOperator()
	Q = gen.New(db)
}
