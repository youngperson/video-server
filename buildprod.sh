#! /bin/bash

# build web and other services
cd ~/code/src/github.com/video-server/api
go build -o ../bin/api

cd ~/code/src/github.com/video-server/scheduler
go build -o ../bin/scheduler

cd ~/code/src/github.com/video-server/streamserver
go build -o ../bin/streamserver

cd ~/code/src/github.com/video-server/web
go build -o ../bin/web


# ========= 目标部署机器的环境设置,编译(目标机器是Linux)=======
#env GOOS=linux GOARCH=amd64 go build -o ../bin/api

#cd ~/code/src/github.com/video-server/scheduler
#env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

#cd ~/code/src/github.com/video-server/streamserver
#env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

#cd ~/code/src/github.com/video-server/web
#env GOOS=linux GOARCH=amd64 go build -o ../bin/web