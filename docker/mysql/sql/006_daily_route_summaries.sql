CREATE TABLE IF NOT EXISTS `daily_route_summaries` (
  `id` int NOT NULL AUTO_INCREMENT,
  `route_id` int NOT NULL,
  `delay_count` int NOT NULL,
  `delay_duration` time NOT NULL,
  `registered_date` date NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`route_id`) REFERENCES `routes` (`id`)
);
