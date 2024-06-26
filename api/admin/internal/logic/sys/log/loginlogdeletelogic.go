package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type LoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogDeleteLogic {
	return LoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LoginLogDelete 删除登录日志
func (l *LoginLogDeleteLogic) LoginLogDelete(req types.DeleteLoginLogReq) (*types.DeleteLoginLogResp, error) {
	_, err := l.svcCtx.LoginLogService.LoginLogDelete(l.ctx, &sysclient.LoginLogDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除登录日志失败")
	}

	return &types.DeleteLoginLogResp{
		Code:    "000000",
		Message: "删除登录日志成功",
	}, nil
}
