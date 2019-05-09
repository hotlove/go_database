CREATE DATABASE `go_project` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci

DROP TABLE IF EXISTS `go_profile`;
CREATE TABLE `go_profile` (
    `id` int(9) NOT NULL,
    `name` varchar(50) DEFAULT NULL,
    `password` varchar(50) DEFAULT NULL,
    `mobile` varchar(11) DEFAULT NULL,
    `email` varchar(40) DEFAULT NULL,
    `area_code` varchar(11) DEFAULT NULL,
    `area_info` varchar(200) DEFAULT NULL,
    `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `avatar` varchar(150) DEFAULT NULL,
    `deleted` int(9) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
