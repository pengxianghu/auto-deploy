#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/auto-deploy
git pull
cd webserver
./webserver &