package model

import (
	"github.com/file_server/common"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var FileConn FilesModel

func init() {
	FileConn = NewFilesModel(sqlx.NewMysql(common.DB_ADDRESS))
}
