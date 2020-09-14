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
  `joinDate` text,
  `location` text,
  PRIMARY KEY (`id`),
  KEY `UserId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# Dump of table `tweets`
# ------------------------------------------------------------

CREATE TABLE `tweets` (
  `id` bigint(20) NOT NULL,
  `userId` bigint(20),
  `userName` text,
  `text` text,
  `created` text,
  `likes` int(11),
  `reTweets` int(11),
  `replyTo` text,
  PRIMARY KEY (`id`),
  KEY `UserId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
