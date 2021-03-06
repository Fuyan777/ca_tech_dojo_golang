package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IController interface {
		Create(c *gin.Context)
	}
	Controller struct {
		UserCreate CreateUsecase
	}
)

func NewUserController(userCreate CreateUsecase) IController {
	return &Controller{UserCreate: userCreate}
}

func (ctrl Controller) Create(c *gin.Context) {
	body := &CreateRequest{}
	if err := c.BindJSON(body); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	res, err := ctrl.UserCreate.Execute(*body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
