// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloud_disk/core/internal/helper"
	"cloud_disk/core/internal/models"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//1.从数据库查询当前用户
	var user = new(models.UserBasic)
	has, err := models.Engine.Where("name = ? and password = ?", req.Name, helper.MD5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	//没有查到用户
	if !has {
		err = errors.New("用户名或密码错误")
		return nil, err
	}
	//生成Token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	return resp, nil
}
