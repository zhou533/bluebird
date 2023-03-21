package handler

import (
	"net/http"

	"bluebird/api/internal/logic"
	"bluebird/api/internal/svc"
	"bluebird/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddSeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeedAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddSeedLogic(r.Context(), svcCtx)
		resp, err := l.AddSeed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
