package sqlc

import (
	"cmd/main.go/sqlc/gen"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/stub"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"jw.lib/sqlx"
)

var Q *gen.Queries
var M *migrate.Migrate

func init() {
	sqlx.Register(sqlx.DefaultSqlDriver, sqlx.DefaultSqlAddr)
	db := sqlx.GetSqlOperator()

	ins, err := stub.WithInstance(db, &stub.Config{})
	if err != nil {
		panic(err)
	}

	M, err = migrate.NewWithDatabaseInstance("file://migration", "postgres", ins)
	if err != nil {
		panic(err)
	}
}
