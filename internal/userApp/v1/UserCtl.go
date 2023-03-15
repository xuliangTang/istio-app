package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

// UserCtl @Controller
type UserCtl struct {
	UserSvc *UserService `inject:"-"`
}

func NewUserCtl() *UserCtl {
	return &UserCtl{}
}

func (this *UserCtl) login(ctx *gin.Context) any {
	kindLoginReq := &KindLoginRequest{}
	athena.Error(ctx.BindJSON(kindLoginReq))

	rsp, err := this.UserSvc.UserLogin(ctx, kindLoginReq)
	athena.Error(err)

	return rsp.Data
}

func (this *UserCtl) Build(athena *athena.Athena) {
	athena.Handle("POST", "/login", this.login)
}
