#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/auto-deploy
git pull https://github.com/pengxianghu/auto-deploy.git
cd webserver
./webserver &