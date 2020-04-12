#! /bin/bash

cp -R ./templates ./bin/

mkdir ./bin/videos  # streamserver中上传到本地的路径

cd bin

nohup ./api &
nohup ./scheduler &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"