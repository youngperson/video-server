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

CREATE TABLE `video_del_rec` (
  `video_id` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```


## 删除视频流程
```
用户->api service->delete video
api service->scheduler->write video deletion record
timer->runner->read video_del_rec->exec->delete video from folder
```

## 流控
```
使用channel实现bucket算法
使用有buffer的channel，buffer的长度就是bucket的个数进行控制
当bucket使用完了说明buffer也使用完了，当还有buffer说明bucket还有
```

## 层次划分
```
每个结构的层次划分清晰，便于后续拆分

web/  :8080   处理前端的操作，把web中的请求处理传递到api中去
templates/    前端的页面、js、image
build.sh      构建前端二进制8080的脚本
api/  :8000   接收web的操作，进行处理
scheduler/  :9001  操作视频删除
streamserver/  :9000  操作视频上传
```

## build.sh
```
sh build.sh
cd ~/code/bin/video_server_web_ui/
./web
浏览器访问http://127.0.0.1:8080/
```