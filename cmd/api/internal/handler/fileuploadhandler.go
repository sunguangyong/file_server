package handler

import (
	"fmt"
	"github.com/file_server/common"
	"log"
	"net/http"

	"github.com/file_server/cmd/api/internal/logic"
	"github.com/file_server/cmd/api/internal/svc"
	"github.com/file_server/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileRequest

		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("jjjjjjjj", r)
			httpx.Error(w, err)
			return
		}

		mp := r.MultipartForm

		fileHeaders, findFile := mp.File["file"]

		if !findFile || len(fileHeaders) == 0 {
			log.Println("file count == 0.")
			err_msg := "ERROR: Not Upload Any File"
			w.Write(common.FormatResponse("", -1, err_msg))
			return
		}

		fmt.Println("11", req.AppName)
		fmt.Println("22", req.TypeName)
		//fmt.Println("33", req.File)

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
