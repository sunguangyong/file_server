package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DirectoryModel = (*customDirectoryModel)(nil)

type (
	// DirectoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDirectoryModel.
	DirectoryModel interface {
		directoryModel
	}

	customDirectoryModel struct {
		*defaultDirectoryModel
	}
)

// NewDirectoryModel returns a model for the database table.
func NewDirectoryModel(conn sqlx.SqlConn) DirectoryModel {
	return &customDirectoryModel{
		defaultDirectoryModel: newDirectoryModel(conn),
	}
}

// 根据文件id 查找路径
/*
WITH RECURSIVE cte_path AS (
  SELECT id, name, prefix, fid, level, uid
  FROM directory
  WHERE fid in (9,10) and uid = 1-- 模糊匹配出来的文件id
  UNION ALL
  SELECT fp.id, fp.name, fp.prefix,fp.fid,fp.level,fp.uid
  FROM directory fp
  JOIN cte_path cte ON cte.prefix = fp.id
	where fp.uid = 1
)
SELECT  DISTINCT *
FROM cte_path
ORDER BY level;
*/
