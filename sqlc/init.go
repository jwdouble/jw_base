package sqlc

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"jw.lib/logx"
	"jw.lib/sqlx"

	"cmd/main.go/sqlc/gen"
)

var Q *gen.Queries
var M *migrate.Migrate

func init() {
	sqlx.Register(sqlx.Driver, sqlx.PGConfigMap)
	db := sqlx.GetSqlOperator()

	// 看文档得看清楚版本和更新日期
	// go mod默认使用v1版本，需要自己确认模块的最新版本是什么
	di, err := postgres.WithInstance(db, &postgres.Config{MigrationsTable: "_migration"})
	if err != nil {
		logx.Error("stub.WithInstance", err)
		return
	}

	M, err = migrate.NewWithDatabaseInstance("file:///home/jw/files/self/jw_base/sqlc/migration", "postgres", di)
	if err != nil {
		logx.Error("migrate.NewWithDatabaseInstance", err)
		return
	}

	M.Log = &Log{}

	fmt.Println(M.Version())

	err = M.Up()
	if err != nil {
		fmt.Println("err-->", err)
	} else {
		fmt.Println("migrate success")
	}
}

func Register() {

}

type Log struct {
}

func (l Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func (l Log) Verbose() bool {
	return true
}

var _ migrate.Logger = (*Log)(nil)
