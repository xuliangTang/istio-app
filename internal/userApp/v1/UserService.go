package v1

import (
	"context"
)

type UserService struct {
	UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) UserLogin(ctx context.Context, req *KindLoginRequest) (*KindLoginResponse, error) {
	return &KindLoginResponse{
		Code:    200,
		Message: "success",
		Data:    &UserModel{UserId: 101, UserName: req.Spec.UserName},
	}, nil
}
