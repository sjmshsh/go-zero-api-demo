package logic

import (
	svc2 "Go-Zero/user-api/internal/svc"
	types2 "Go-Zero/user-api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc2.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types2.UserInfoReq) (resp *types2.UserInfoResp, err error) {
	return nil, err
}
