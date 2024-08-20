package logic

import (
	"context"

	ierrors "go-url-shortener/errors"
	"go-url-shortener/internal/svc"
	"go-url-shortener/internal/types"
	"go-url-shortener/internal/utils"
	"go-url-shortener/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenRequest) (resp *types.ShortenResponse, err error) {
	// test connection
	// if the url is unreachable, there is no need to proceed
	if !utils.TestConnection(req.OriginalURL) {
		return nil, ierrors.ErrInvalidURL
	}

	// check if the url is already shortened and present in the database
	// only proceed if the following lookup returns ErrNotFound
	res, err := l.svcCtx.URLMapModel.FindOneByMd5(l.ctx, utils.GenerateMD5(req.OriginalURL))
	if err != model.ErrNotFound {
		if err == nil {
			return &types.ShortenResponse{
				ShortURL: res.ShortUrl,
			}, nil
		} else {
			logx.Errorw("failed to find url by md5", logx.LogField{Key: "error", Value: err.Error()})
			return nil, err
		}
	}

	return
}
