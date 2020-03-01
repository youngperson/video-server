# video-server
## 数据库设计
```
mysql5.8+

CREATE TABLE `sessions` (
  `session_id` tinytext NOT NULL,
  `TTL` tinytext,
  `login_name` text,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(64) DEFAULT NULL,
  `pwd` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name` (`login_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `video_info` (
  `id` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `author_id` int(10) DEFAULT NULL,
  `name` text,
  `display_ctime` text,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `comments` (
  `id` varchar(64) NOT NULL DEFAULT '',
  `video_id` varchar(64) DEFAULT NULL,
  `author_id` int(10) DEFAULT NULL,
  `content` text,
  `time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```
