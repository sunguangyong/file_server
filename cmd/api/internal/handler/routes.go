// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/file_server/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload/:AppName/:TypeName",
				Handler: FileUploadHandler(serverCtx),
			},
		},
	)
}
