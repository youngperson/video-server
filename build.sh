#! /bin/bash

# build web UI
cd ~/code/src/github.com/video-server/web
go install
# 把目录创建出来
cp ~/code/bin/web ~/code/bin/video_server_web_ui/web
cp -R  ~/code/src/github.com/video-server/templates ~/code/bin/video_server_web_ui/
