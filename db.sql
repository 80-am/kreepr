# Dump of table followers
# ------------------------------------------------------------

CREATE TABLE `followers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `userName` varchar(100) NOT NULL,
  `status` text,
  `followDate` datetime NOT NULL,
  `unfollowDate` datetime DEFAULT NULL,
  `lastAction` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `UserId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
