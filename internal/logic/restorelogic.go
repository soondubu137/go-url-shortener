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
	shortURL := req.ShortURL
	// query for the original url from the database
	res, err := l.svcCtx.URLMapModel.FindOneByShortUrl(l.ctx, shortURL)
	if err != nil {
		logx.Errorw("failed to find url by short url", logx.LogField{Key: "error", Value: err.Error()})
		return nil, err
	}

	return &types.RestoreResponse{
		OriginalURL: res.OriginalUrl,
	}, nil
}
