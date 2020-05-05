package user

import (
	"game/app/pkg/jwt"
)

type (
	CreateRequest struct {
		Name string `json:"name"`
	}
	CreateResponse struct {
		Token string `json:"token"`
	}
	CreateUsecase interface {
		Execute(req CreateRequest) (*CreateResponse, error)
	}
	CreateInteractor struct {
		UseRepo Repository
	}
)

func NewCreateInteractor(useRepo Repository) CreateUsecase {
	return &CreateInteractor{UseRepo: useRepo}
}

func (c *CreateInteractor) Execute(req CreateRequest) (*CreateResponse, error) {
	token, err := jwt.Create(req.Name)
	if err != nil {
		return nil, err
	}
	user := &User{
		Name:  req.Name,
		Token: token,
	}
	if err := c.UseRepo.Create(*user); err != nil {
		return nil, err
	}
	return &CreateResponse{Token: token}, nil
}
