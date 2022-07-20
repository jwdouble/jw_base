package service

import (
	"context"

	"cmd/main.go/orm"
	"cmd/main.go/orm/gen"
)

func CreateText(ctx context.Context, arg gen.CreateTextParams) error {
	err := orm.Q.CreateText(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
