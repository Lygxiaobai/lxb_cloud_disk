// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloud_disk/core/internal/define"
	"cloud_disk/core/internal/models"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	list := []*types.UserFile{}
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = define.Page
	}
	offset := (page - 1) * size
	l.svcCtx.Engine.ShowSQL(true)
	//1.联表查询属于UserIdentity用户的所有文件信息 注意要结合父级id查询
	err = l.svcCtx.Engine.Table("user_repository").Where("user_identity=? AND parent_id =?", userIdentity, req.Id).
		Select("user_repository.parent_id,user_repository.identity,user_repository.repository_identity,user_repository.name,user_repository.ext,repository_pool.size,repository_pool.path").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Limit(size, offset).
		Where("user_repository.deleted_at IS NULL").
		Find(&list)
	if err != nil {
		return nil, err
	}
	count, err := l.svcCtx.Engine.Where("user_identity=? AND parent_id =?", userIdentity, req.Id).Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	//2.返回数据
	resp = &types.UserFileListResponse{
		list,
		count,
	}
	return
}
