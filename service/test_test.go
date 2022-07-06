package service

import (
	"cmd/main.go/tutorial"
	"context"
	"fmt"
	"testing"
)

func TestCreateText(t *testing.T) {
	err := CreateText(context.Background(), tutorial.CreateTextParams{
		ID: "1",
		T:  "hha",
	})
	if err != nil {
		fmt.Println(err)
	}
}
