// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-url-shortener/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:short_url",
				Handler: RestoreHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shorten",
				Handler: ShortenHandler(serverCtx),
			},
		},
	)
}
