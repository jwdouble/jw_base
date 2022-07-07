package service

import (
	"cmd/main.go/sqlc/gen"
	"context"
	"fmt"
	"testing"
)

func TestCreateText(t *testing.T) {
	err := CreateText(context.Background(), gen.CreateTextParams{
		ID: "2",
		T:  "hha",
	})
	if err != nil {
		fmt.Println(err)
	}
}
