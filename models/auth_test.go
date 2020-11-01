package models

import "testing"

func TestCreteTable(t *testing.T) {
	err := CreteTable()
	t.Log(err)
}
func TestAddAuth(t *testing.T) {
	err := AddAuth("admin", "admin")
	t.Log(err)
}
