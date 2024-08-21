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
	if !utils.CanConnect(req.OriginalURL) {
		return nil, ierrors.ErrInvalidURL
	}

	// check if the url has already been shortened and exists in the database
	// only proceed if the following lookup returns ErrNotFound
	md5 := utils.GenerateMD5(req.OriginalURL)
	res, err := l.svcCtx.URLMapModel.FindOneByMd5(l.ctx, md5)
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

	// check if the url is itself a shortened url, we don't accept such urls
	baseURL, err := utils.GetBaseURL(req.OriginalURL)
	if err != nil {
		logx.Errorw("failed to get base url", logx.LogField{Key: "error", Value: err.Error()})
		return nil, err
	}
	_, err = l.svcCtx.URLMapModel.FindOneByShortUrl(l.ctx, baseURL)
	if err != model.ErrNotFound {
		if err == nil {
			return nil, ierrors.ErrURLAlreadyShortened
		} else {
			logx.Errorw("failed to find url by short url", logx.LogField{Key: "error", Value: err.Error()})
			return nil, err
		}
	}

	// generate a new id
	id, err := l.svcCtx.IdGenerator.Generate(l.ctx, "dflt")
	if err != nil {
		logx.Errorw("failed to generate id", logx.LogField{Key: "error", Value: err.Error()})
		return nil, err
	}
	shortURL := utils.EncodeBase62(id, l.svcCtx.Config.CypherKey)

	// insert the pair of original URL and shortened URL into the database
	_, err = l.svcCtx.URLMapModel.Insert(l.ctx, &model.UrlMap{
		Md5:         md5,
		OriginalUrl: req.OriginalURL,
		ShortUrl:    shortURL,
	})
	if err != nil {
		logx.Errorw("failed to insert url map", logx.LogField{Key: "error", Value: err.Error()})
		return nil, err
	}

	// return the shortened URL
	return &types.ShortenResponse{
		ShortURL: shortURL,
	}, nil
}
