package logic

import (
	"context"

	"go-url-shortener/internal/svc"
	"go-url-shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreLogic {
	return &RestoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestoreLogic) Restore(req *types.RestoreRequest) (resp *types.RestoreResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
