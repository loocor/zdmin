package handler

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/couponhistory"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CouponHistoryAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCouponHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCouponHistoryAddLogic(r.Context(), ctx)
		resp, err := l.CouponHistoryAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
