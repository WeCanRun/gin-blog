package model

import (
	"context"
	"testing"
)

func TestCreteTable(t *testing.T) {
	err := CreteTable(context.Background())
	t.Log(err)
}
func TestAddAuth(t *testing.T) {
	err := AddAuth(context.Background(), "admin", "admin")
	t.Log(err)
}
