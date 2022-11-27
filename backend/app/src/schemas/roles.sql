CREATE TABLE `roles` (
  `role_id` int NOT NULL AUTO_INCREMENT,
  `role_icon` varchar(127) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_name` varchar(30) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
