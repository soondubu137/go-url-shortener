package handler

import (
	"net/http"

	"go-url-shortener/internal/logic"
	"go-url-shortener/internal/svc"
	"go-url-shortener/internal/types"
	"go-url-shortener/model"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/go-playground/validator/v10"
)

func RestoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RestoreRequest
		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// Parameter validation
		if err := validate.Struct(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRestoreLogic(r.Context(), svcCtx)
		resp, err := l.Restore(&req)
		if err != nil {
			if err == model.ErrNotFound {
				httpx.ErrorCtx(r.Context(), w, model.ErrNotFound)
			} else {
				httpx.ErrorCtx(r.Context(), w, err)
			}
		} else {
			http.Redirect(w, r, resp.OriginalURL, http.StatusFound)
		}
	}
}
