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

type GetFileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileInfoLogic) GetFileInfo(req *types.FileInfoRequest) (resp *types.FileOnfoResponse, err error) {
	// todo: add your logic here and delete this line

	resp = &types.FileOnfoResponse{}
	file, err := model.FileConn.FindOne(l.ctx, req.FileId)

	if err != nil {
		return
	}

	if file == nil {
		return nil, xerr.NewErr(201, errors.New("没有此文件"))
	}

	resp.DiskPath = file.DiskPath
	resp.DownloadPath = file.DownloadPath

	return
}
