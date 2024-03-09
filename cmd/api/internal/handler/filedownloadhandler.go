package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/file_server/cmd/api/internal/logic"
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

		l := logic.NewFileDownloadLogic(r.Context(), svcCtx)

		file, err := l.FileDownload(&req)

		if err != nil {
			return
		}

		diskPath := file.DiskPath

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.OrigFileName))

		if !fileutil.FileExist(diskPath) {
			xerr.NewErr(201, errors.New("ERROR: file Not Found"))
			return
		}

		http.ServeFile(w, r, diskPath)
	}
}
