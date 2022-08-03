package service

import (
	"context"

	"jw_base/orm"
	"jw_base/orm/gen"
)

func CreateText(ctx context.Context, arg gen.CreateTextParams) error {
	err := orm.Q.CreateText(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
