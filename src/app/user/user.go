package user

import "game/app/pkg/gorm"

type User struct {
	gorm.BaseModel
	Name  string `json:"name"`
	Token string `json:token`
}
