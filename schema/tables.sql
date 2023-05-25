-- 用户表 --
CREATE TABLE `user`
(
    `id`         BIGINT(11) NOT NULL AUTO_INCREMENT,
    `username`   VARCHAR(32) NOT NULL COMMENT '用户名',
    `password`   VARCHAR(32) NOT NULL COMMENT '用户密码',
    `created_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `_u` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';