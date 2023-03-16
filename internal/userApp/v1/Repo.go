package v1

import (
	"fmt"
	"math/rand"
)

type IUserRepo interface {
	FindByID(model *UserModel) error
	New(model *UserModel) error
	Update(model *UserModel) error
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (this *UserRepo) FindByID(model *UserModel) error {
	if model.UserName == "txl" {
		model.UserPwd = "123"
	} else if model.UserName == "lisi" {
		model.UserPwd = "234"
	} else {
		return fmt.Errorf("no such user %s", model.UserName)
	}
	return nil
}

func (this *UserRepo) Update(model *UserModel) error {
	return nil
}

func (this *UserRepo) New(model *UserModel) error {
	model.UserId = int32(rand.Intn(100))

	if model.UserName == "txl" {
		return fmt.Errorf("用户名已经存在")
	}

	return nil
}
