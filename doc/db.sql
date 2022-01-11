CREATE TABLE demo_user (
    id int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    sex tinyint NULL DEFAULT 0 COMMENT '性别:0男,1女',
    age tinyint NULL DEFAULT 0 COMMENT '年龄',
    created_at timestamp NULL DEFAULT NULL COMMENT '创建时间',
    updated_at timestamp NULL DEFAULT NULL COMMENT '更新时间',
    deleted_at timestamp NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;