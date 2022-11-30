package handler

import (
	logic2 "Go-Zero/user-api/internal/logic"
	svc2 "Go-Zero/user-api/internal/svc"
	types2 "Go-Zero/user-api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 参数解析
		var req types2.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 具体业务写在这里
		l := logic2.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
