CREATE TABLE `task_status` (
  `task_state_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `task_id` int NOT NULL,
  `task_flag_id` int NOT NULL,
  `progress` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`task_state_id`),
  KEY `uses_index` (`user_id`),
  KEY `tasks_index` (`task_id`),
  KEY `task_flags_index` (`task_flag_id`),
  CONSTRAINT `task_status_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `task_status_ibfk_2` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`task_id`),
  CONSTRAINT `task_status_ibfk_3` FOREIGN KEY (`task_flag_id`) REFERENCES `task_flags` (`task_flag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
