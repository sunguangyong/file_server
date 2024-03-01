CREATE TABLE `file_path` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `path_name` varchar(256) NOT NULL DEFAULT '' COMMENT '路径名称',
     `up_path_id` int NOT NULL DEFAULT '0' COMMENT '上行路径id',
     `file_id` int NOT NULL DEFAULT '0' COMMENT '文件id',
     `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='路径表'