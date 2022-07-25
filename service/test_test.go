package service

import (
	"context"
	"testing"

	"cmd/main.go/orm/gen"
)

func TestCreateText(t *testing.T) {
	err := CreateText(context.Background(), gen.CreateTextParams{
		ID: "2",
		T:  "hha",
	})
	if err != nil {
		t.Log(err)
	}
}
