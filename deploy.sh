#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/auto-deploy
git pull git@github.com:pengxianghu/auto-deploy.git
cd webserver
./webserver &