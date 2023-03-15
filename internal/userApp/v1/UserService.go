package v1

import (
	"context"
	"fmt"
)

type UserService struct {
	UnimplementedUserServiceServer
	UserRepo IUserRepo `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) UserLogin(ctx context.Context, req *KindLoginRequest) (*KindLoginResponse, error) {
	userModel := &UserModel{
		UserName: req.Spec.UserName,
	}

	err := this.UserRepo.FindByID(userModel)
	if err != nil {
		return nil, err
	}

	if userModel.UserPwd != req.Spec.UserPwd {
		return nil, fmt.Errorf("error username or password")
	}

	return &KindLoginResponse{
		Code:    200,
		Message: "success",
		Data:    &UserModel{UserId: 101, UserName: userModel.UserName},
	}, nil
}
