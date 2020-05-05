package main

import (
	"game/app/pkg/mysql"
	"game/app/server"
	"game/app/user"
	"log"
)

func main() {
	mysqlDriver := mysql.NewSQLDriver()
	if err := mysqlDriver.SqlConn(); err != nil {
		log.Fatalln(err)
	}
	if err := mysqlDriver.GetGorm().AutoMigrate(&user.User{}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Printf("info: create user table\n")
	userGateway := user.NewGateway(mysqlDriver)
	userCreateInteractor := user.NewCreateInteractor(userGateway)
	userCtrl := user.NewUserController(userCreateInteractor)

	s := server.NewServer(userCtrl)
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}
