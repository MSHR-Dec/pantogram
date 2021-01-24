CREATE TABLE IF NOT EXISTS `regions` (
  `id` tinyint(3) unsigned NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `regions` VALUES
  (1,'北海道地方'),
  (2,'東北地方'),
  (3,'関東地方'),
  (4,'中部地方'),
  (5,'近畿地方'),
  (6,'中国地方'),
  (7,'四国地方'),
  (8,'九州地方');
