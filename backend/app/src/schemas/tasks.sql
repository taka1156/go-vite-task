CREATE TABLE `tasks` (
  `task_id` int NOT NULL AUTO_INCREMENT,
  `task_name` varchar(127) COLLATE utf8mb4_general_ci NOT NULL,
  `link` varchar(127) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `level` int NOT NULL DEFAULT '1',
  `description` varchar(4000) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_review` tinyint(1) NOT NULL DEFAULT '0',
  `is_all_required` tinyint(1) NOT NULL DEFAULT '0',
  `category_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`task_id`),
  KEY `categories_index` (`category_id`),
  CONSTRAINT `tasks_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
