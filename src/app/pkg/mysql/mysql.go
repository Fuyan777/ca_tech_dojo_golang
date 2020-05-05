package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SQLDriver interface {
	GetGorm() *gorm.DB
	SqlConn() error
}

type Driver struct {
	Gorm *gorm.DB
}

func NewSQLDriver() SQLDriver {
	return &Driver{}
}

func (d *Driver) SqlConn() (err error) {
	d.Gorm, err = gorm.Open("mysql", "docker:docker@tcp(mysql:3306)/game?parseTime=true")
	if err != nil {
		return
	}
	return
}

func (d *Driver) GetGorm() *gorm.DB {
	return d.Gorm
}
