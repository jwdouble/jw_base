package service

import (
	"cmd/main.go/sqlc"
	"cmd/main.go/sqlc/gen"
	"context"
)

func CreateText(ctx context.Context, arg gen.CreateTextParams) error {
	err := sqlc.Q.CreateText(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
