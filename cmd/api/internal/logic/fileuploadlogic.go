package logic

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	common "github.com/file_server/common/constant"
	"github.com/file_server/common/fileutil"
	"github.com/file_server/common/xerr"
	"github.com/file_server/model"

	"github.com/file_server/cmd/api/internal/svc"
	"github.com/file_server/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest, mp *multipart.Form) (resp *types.FileUploadResponse,
	err error) {

	resp = &types.FileUploadResponse{
		Data: make([]types.FileUpload, 0),
	}

	fileHeaders, findFile := mp.File["file"]

	if !findFile || len(fileHeaders) == 0 {
		return nil, xerr.NewErr(201, errors.New("没有此文件"))
	}

	for _, fh := range fileHeaders {
		data, err := l.save(req, fh)
		if err != nil {
			continue
		}
		resp.Data = append(resp.Data, data)
	}

	return
}

func (l *FileUploadLogic) save(req *types.FileUploadRequest, fh *multipart.FileHeader) (resp types.FileUpload,
	err error) {

	size := fh.Size
	origFileName := fh.Filename
	fileName := fileutil.GenFileName(origFileName)
	appName := req.AppName
	typeName := req.TypeName

	downloadPath := fileutil.DoloadFilePath(appName, typeName, url.QueryEscape(fileName)) // 将字符地址转为安全地址  防止特殊符号在服务器端无法获得正确的参数值
	absolutePath := fileutil.AbsoluteFilePath(appName, typeName, url.QueryEscape(fileName))

	diskDir := fileutil.CurrentDirPath(appName, typeName)

	err = fileutil.CheckOrMakedir(diskDir)

	if err != nil {
		return resp, xerr.NewErr(201, errors.New(err.Error()))
	}

	diskPath := fmt.Sprintf("%s/%s", diskDir, fileName)
	err = fileutil.SaveFile(diskPath, fh)

	if err != nil {
		return resp, xerr.NewErr(201, errors.New(err.Error()))
	}

	files := &model.Files{}
	files.FileName = fileName
	files.DownloadPath = common.STATIC_URL + downloadPath
	files.AbsolutePath = absolutePath
	files.DiskPath = diskPath
	files.FileSize = size
	files.FileType = typeName
	files.AppName = appName
	files.OrigFileName = origFileName
	files.HostName, _ = os.Hostname()
	files.StorageIp = common.STORAGE_IP
	files.CreateTime = time.Now()
	files.UpdateTime = time.Now()

	result, err := model.FileConn.Insert(l.ctx, files)

	if err != nil {
		return resp, xerr.NewErr(201, errors.New(err.Error()))
	}

	fileId, err := result.LastInsertId()

	if err != nil {
		return resp, xerr.NewErr(201, errors.New(err.Error()))
	}

	resp = types.FileUpload{
		FileId:       fileId,
		DownloadPath: files.DownloadPath,
		DiskPath:     files.DiskPath,
	}

	return
}
