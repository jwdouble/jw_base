package service

import (
	"cmd/main.go/tutorial"
	"context"
	"jw.lib/sqlx"
)

func CreateText(ctx context.Context, arg tutorial.CreateTextParams) error {
	sqlx.Register(sqlx.DefaultSqlDriver, sqlx.DefaultSqlAddr)

	DBQ := tutorial.New(sqlx.GetSqlOperator())
	err := DBQ.CreateText(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
