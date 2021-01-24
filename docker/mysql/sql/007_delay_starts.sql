CREATE TABLE IF NOT EXISTS `delay_starts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `route_id` int NOT NULL,
  `start_time` time NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`route_id`) REFERENCES `routes` (`id`)
);
