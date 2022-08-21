package model

import (
	"context"
	otgorm "github.com/EDDYCJY/opentracing-gorm"
	"log"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreteTable(ctx context.Context) error {
	return otgorm.WithContext(ctx, db).CreateTable(&Auth{}).Error
}

func AddAuth(ctx context.Context, username, password string) error {
	return otgorm.WithContext(ctx, db).Create(&Auth{
		Username: username,
		Password: password,
	}).Error
}

func CheckAuth(ctx context.Context, username, password string) bool {
	var auth Auth
	err := otgorm.WithContext(ctx, db).Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		log.Printf("CheckAuth fail, %v", err)
	}
	return auth.ID > 0
}
