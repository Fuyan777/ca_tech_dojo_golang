package server

import (
	"game/app/user"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router   *gin.Engine
	UserCtrl user.IController
}

func NewServer(userCtrl user.IController) *Server {
	result := &Server{}
	result.Router = gin.Default()
	result.UserCtrl = userCtrl
	_ = result.SetupRouting()
	return result
}

func (s *Server) SetupRouting() error {
	router := gin.Default()
	router.POST("/user/create", s.UserCtrl.Create)

	s.Router = router

	return nil
}

func (s *Server) Run() error {
	return s.Router.Run(":8080")
}
