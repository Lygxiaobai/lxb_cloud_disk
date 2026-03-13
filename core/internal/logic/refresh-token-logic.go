// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloud_disk/core/internal/define"
	"cloud_disk/core/internal/helper"
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest, authorization string) (resp *types.RefreshTokenResponse, err error) {
	//这里传的应该使refreshToken
	//根据token 获取userClaim
	uc, err := helper.AnalyzeToken(authorization)
	if err != nil {
		return nil, err
	}
	if uc == nil {
		return nil, errors.New("invalid authorization")
	}
	//根据uc来生成新的token
	token, err := helper.GenerateToken(uc.ID, uc.Identity, uc.Name, define.TokenExpireTime)
	if err != nil {
		return nil, err
	}
	refreshToken, err := helper.GenerateToken(uc.ID, uc.Identity, uc.Name, define.TokenExpireTime)
	if err != nil {
		return nil, err
	}
	resp = &types.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return
}
