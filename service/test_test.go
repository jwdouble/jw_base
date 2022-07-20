package service

import (
	"context"
	"fmt"
	"testing"

	"cmd/main.go/orm/gen"
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
