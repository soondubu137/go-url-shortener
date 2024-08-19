CREATE TABLE url_map (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` VARCHAR(64) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL DEFAULT FALSE,
    `original_url` VARCHAR(512) NOT NULL,
    `short_url` VARCHAR(16) NOT NULL,
    `md5` CHAR(32) NOT NULL,
    PRIMARY KEY (`id`),
    INDEX(`is_deleted`),
    UNIQUE(`md5`),
    UNIQUE(`short_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;