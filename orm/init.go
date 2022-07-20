package orm

import (
	"embed"

	"jw.lib/migrate"
	"jw.lib/sqlx"

	"cmd/main.go/orm/gen"
)

//go:embed migration/*.sql
var migration embed.FS

var Q *gen.Queries

func init() {
	sqlx.Register(sqlx.Driver, sqlx.PGConfigMap)
	migrate.NewWithFs(migration, "migration", sqlx.GetSqlOperator())
}

func Register() {

}

//type Log struct {
//}
//
//func (l Log) Printf(format string, v ...interface{}) {
//	fmt.Printf(format+"\n", v...)
//}
//
//func (l Log) Verbose() bool {
//	return true
//}
//
//var _ migrate.Logger = (*Log)(nil)
