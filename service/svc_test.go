package service

import (
	"context"
	"testing"

	"jw.lib/conf"
)

var ctx = context.Background()
var baseSvc = NewBaseService()

func Test_mass(t *testing.T) {
	t.Log(conf.SERVER_PORT.Value(":10000"))
	t.Log(conf.Get("app.network.tls_enable"))
}
