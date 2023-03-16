package v1

import (
	"context"
	"fmt"
)

const (
	UserStatusLoginSuccess    = 200
	UserStatusRegisterSuccess = 1200
)
const (
	MessageSuccess = "success"
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
		Code:    UserStatusLoginSuccess,
		Message: "success",
		Data:    &UserModel{UserId: 101, UserName: userModel.UserName},
	}, nil
}

func (this *UserService) UserRegister(ctx context.Context, req *KindRegisterRequest) (*KindRegisterResponse, error) {
	user := &UserModel{
		UserName: req.Spec.UserName,
		UserPwd:  req.Spec.UserPass,
	}

	err := this.UserRepo.New(user)
	if err != nil {
		return nil, err
	}

	user.UserPwd = ""

	return &KindRegisterResponse{
		Code:    UserStatusRegisterSuccess,
		Message: MessageSuccess,
		Data:    user,
	}, nil
}
