-- ink.s_apps definition

CREATE TABLE `s_apps` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `secret` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `s_apps_secret_key` (`secret`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- ink.s_auth_tokens definition

CREATE TABLE `s_auth_tokens` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `access_token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `refresh_token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_id` int NOT NULL,
  `account_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `s_auth_tokens_access_token_key` (`access_token`),
  UNIQUE KEY `s_auth_tokens_refresh_token_key` (`refresh_token`),
  KEY `s_auth_tokens_app_id_account_id_idx` (`app_id`,`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- ink.s_employees definition

CREATE TABLE `s_employees` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `email` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  `super` tinyint(1) NOT NULL DEFAULT '0',
  `phone` char(11) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `s_employees_email_key` (`email`),
  UNIQUE KEY `s_employees_phone_key` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- ink.s_roles definition

CREATE TABLE `s_roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `s_roles_app_id_name_key` (`app_id`,`name`),
  KEY `s_roles_app_id_idx` (`app_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- ink.s_users definition

CREATE TABLE `s_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `email` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `s_users_email_key` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;