package gorm

import "time"

type BaseModel struct {
	ID       uint      `gorm:"primary_key" json:"-"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
	DeleteAt time.Time `json:"-"`
}
