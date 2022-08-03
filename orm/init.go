package orm

import (
	"embed"

	"jw.lib/migrate"
	"jw.lib/sqlx"

	"jw_base/orm/gen"
)

//go:embed migration/*.sql
var migration embed.FS

var Q *gen.Queries

func init() {
	sqlx.Register(sqlx.Driver, sqlx.PGConfigMap)
	migrate.NewWithFs(migration, "migration", sqlx.GetSqlOperator())
}
