package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FilePathModel = (*customFilePathModel)(nil)

type (
	// FilePathModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilePathModel.
	FilePathModel interface {
		filePathModel
	}

	customFilePathModel struct {
		*defaultFilePathModel
	}
)

// NewFilePathModel returns a model for the database table.
func NewFilePathModel(conn sqlx.SqlConn) FilePathModel {
	return &customFilePathModel{
		defaultFilePathModel: newFilePathModel(conn),
	}
}
