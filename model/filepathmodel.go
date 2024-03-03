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

// 查看某个用户的所有目录层级
/*
WITH RECURSIVE PathHierarchy AS (
SELECT id, path_name, up_path_id, file_id, 0 AS level
FROM file_path
WHERE user_id = 1  AND up_path_id = 0

UNION ALL

SELECT fp.id, fp.path_name, fp.up_path_id, fp.file_id, ph.level + 1
FROM file_path fp
JOIN PathHierarchy ph ON fp.up_path_id = ph.id
)

SELECT id, path_name, level, file_id
FROM PathHierarchy;

*/

// 查看包含某些文件id的所有上层目录
/*
WITH RECURSIVE cte_path AS (
  SELECT id, path_name, up_path_id, file_id
  FROM file_path
  WHERE file_id in (9,10) -- 模糊匹配出来的文件id
  UNION ALL
  SELECT fp.id, fp.path_name, fp.up_path_id,fp.file_id
  FROM file_path fp
  JOIN cte_path cte ON cte.up_path_id = fp.id
)
SELECT  DISTINCT *
FROM cte_path
ORDER BY up_path_id;
*/
