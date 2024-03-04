package handler

import (
	"errors"
	"net/http"
	"strings"

	common "github.com/file_server/common/constant"
	"github.com/file_server/common/fileutil"
	"github.com/file_server/common/xerr"

	"github.com/file_server/cmd/api/internal/svc"
	"github.com/file_server/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		diskPath := strings.Join([]string{common.STATIC_DIR, req.AppName, req.TypeName, req.FileName}, "/")
		if !fileutil.FileExist(diskPath) {
			xerr.NewErr(201, errors.New("ERROR: file Not Found"))
			return
		}
		http.ServeFile(w, r, diskPath)
	}
}
