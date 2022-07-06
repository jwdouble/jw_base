// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package tutorial

import (
	"context"
)

const createText = `-- name: CreateText :exec
insert into test_text (id, t) values ($1,$2)
`

type CreateTextParams struct {
	ID string
	T  string
}

func (q *Queries) CreateText(ctx context.Context, arg CreateTextParams) error {
	_, err := q.db.ExecContext(ctx, createText, arg.ID, arg.T)
	return err
}