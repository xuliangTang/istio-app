package main

import (
	"context"
	v1 "istio-app/internal/userApp/v1"
	"log"
)

func main() {
	svc := &v1.UserService{}
	loginRet, _ := svc.UserLogin(context.Background(), &v1.KindLoginRequest{
		Spec: &v1.UserLoginModel{
			UserName: "txl",
			UserPwd:  "123456",
		},
	})

	log.Println(loginRet)
}
