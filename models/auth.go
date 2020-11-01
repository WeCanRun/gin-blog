package models

import (
	"log"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreteTable() error {
	return db.CreateTable(&Auth{}).Error
}

func AddAuth(username, password string) error  {
	return db.Create(&Auth{
		Username: username,
		Password: password,
	}).Error
}

func CheckAuth(username, password string) bool {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		log.Printf("CheckAuth fail, %v", err)
	}
	return auth.ID > 0
}
