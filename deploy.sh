#! /bin/sh

kill -9 $(pgrep webserver)
git pull https://github.com/pengxianghu/auto-deploy.git
cd ~/auto-deploy/webserver
./webserver &