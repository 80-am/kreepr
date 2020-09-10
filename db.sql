# Dump of table subjects
# ------------------------------------------------------------

CREATE TABLE `subjects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20),
  `name` text,
  `userName` text,
  `followers` int(11),
  `friends` int(11),
  `tweets` bigint(20),
  `location` text,
  PRIMARY KEY (`id`),
  KEY `UserId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
