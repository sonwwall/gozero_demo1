package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_demo1/api_study/user/api_jwt/internal/logic"
	"gozero_demo1/api_study/user/api_jwt/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
