package sqlc

import (
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"

	"cmd/main.go/sqlc/gen"
)

var Q *gen.Queries
var M *migrate.Migrate

func init() {
	//sqlx.Register(sqlx.DefaultSqlDriver, sqlx.DefaultSqlAddr)
	//db := sqlx.GetSqlOperator()
	//
	//ins, err := stub.WithInstance(db, &stub.Config{})
	//if err != nil {
	//	logx.Error("stub.WithInstance",err)
	//	return
	//}
	//
	//M, err = migrate.NewWithDatabaseInstance("file://sqlc/migration", "postgres", ins)
	//if err != nil {
	//	logx.Error("migrate.NewWithDatabaseInstance",err)
	//	return
	//}
	//
	//err = M.Up()
	//if err != nil {
	//	fmt.Println("err-->", err)
	//} else {
	//	fmt.Println("migrate success")
	//}
}

func Register() {

}
