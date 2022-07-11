package sqlc

import (
	"fmt"
	"testing"
)

func Test_migrate(t *testing.T) {
	err := M.Up()
	if err != nil {
		fmt.Println(err)
	}
}
