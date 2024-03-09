package logic

import (
	"context"

	"github.com/file_server/common/xerr"
	"github.com/file_server/model"
	"github.com/pkg/errors"

	"github.com/file_server/cmd/api/internal/svc"
	"github.com/file_server/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDownloadLogic) FileDownload(req *types.FileDownloadRequest) (file *model.Files, err error) {
	file, err = model.FileConn.FindOne(l.ctx, req.FileId)

	if err != nil {
		return
	}

	if file == nil {
		return nil, xerr.NewErr(201, errors.New("没有此文件"))
	}

	return
}
