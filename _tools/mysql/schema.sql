CREATE TABLE `user`
(
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'user identifier',
	`name` VARCHAR(20) NOT NULL COMMENT 'user name',
	`password` VARCHAR(80) NOT NULL COMMENT 'hashed password',
	`role` VARCHAR(80) NOT NULL COMMENT 'user role',
	`created` DATETIME(6) NOT NULL COMMENT 'created time',
	`modified` DATETIME(6) NOT NULL COMMENT 'modified time',
	PRIMARY KEY (`id`),
	UNIQUE KEY `uix_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user table';

CREATE TABLE `task`
(
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'task identifier',
	`title` VARCHAR(120) NOT NULL COMMENT 'task title',
	`status` VARCHAR(20) NOT NULL COMMENT 'task status',
	`created` DATETIME(6) NOT NULL COMMENT 'created time',
	`modified` DATETIME(6) NOT NULL COMMENT 'modified time',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='task table';
