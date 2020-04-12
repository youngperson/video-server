# video-server
## 部署流程
```
1、修改连接数据库的地方,api/dbops/conn.go、scheduler/dbops/conn.go
2、导入initdb.sql
3、修改bin/conf.json，默认为127.0.0.0
4、执行 sh buildprod.sh
5、执行 sh deploy.sh
6、访问浏览器 127.0.0.0:8080
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
每个结构的层次划分清晰，便于后续拆分。
像内部的后台系统可以按照这种拆分的思想去做

web/  :8080   处理前端的操作，把web中的请求处理传递到api中去
templates/    前端的页面、js、image
build.sh      构建前端二进制8080的脚本
api/  :8000   接收web的操作，进行处理(部署到内网即可,不需要对用户暴露)
scheduler/  :9001  操作视频删除(部署到内网即可,不需要对用户暴露)
streamserver/  :9000  操作上传视频、下载视频(部署到内网即可,不需要对用户暴露)
vendor/       把公共的config放进去,省得到处copy
```

## build.sh
```
sh build.sh
cd ~/code/bin/video_server_web_ui/
./web
浏览器访问http://127.0.0.1:8080/
```