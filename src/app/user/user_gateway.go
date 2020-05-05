package user

import (
	"fmt"
	"game/app/pkg/mysql"
)

type (
	Repository interface {
		Create(user User) error
	}
	Gateway struct {
		mysql.SQLDriver
	}
)

func NewGateway(driver mysql.SQLDriver) Repository {
	return &Gateway{SQLDriver: driver}
}

func (g *Gateway) Create(user User) error {
	if err := g.SQLDriver.GetGorm().Create(&user).Error; err != nil {
		return fmt.Errorf("call gorm create user %w", err)
	}
	return nil
}
