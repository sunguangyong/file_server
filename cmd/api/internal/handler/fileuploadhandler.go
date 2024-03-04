package handler

import (
	"net/http"

	"github.com/file_server/common/result"

	"github.com/file_server/cmd/api/internal/logic"
	"github.com/file_server/cmd/api/internal/svc"
	"github.com/file_server/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		mp := r.MultipartForm
		//fh := r.MultipartFileHeader

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req, mp)

		result.HttpResult(r, w, req, resp, err)

	}
}
