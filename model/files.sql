CREATE TABLE `files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` INT(10) DEFAULT '0' COMMENT '用户id',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
  `host_name` varchar(256) NOT NULL DEFAULT '' COMMENT '主机名',
  `disk_path` varchar(256) NOT NULL DEFAULT '' COMMENT '磁盘路径',
  `download_path` varchar(256) NOT NULL DEFAULT '' COMMENT '下载路径',
  `absolute_path` varchar(256) NOT NULL DEFAULT '' COMMENT '名称',
  `app_name` varchar(256) NOT NULL DEFAULT '' COMMENT 'app名称',
  `file_type` varchar(256) NOT NULL DEFAULT '' COMMENT '文件类型',
  `storage_ip` varchar(256) NOT NULL DEFAULT '' COMMENT '存储ip',
  `file_size` int NOT NULL DEFAULT 0 COMMENT '文件大小',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件表'

