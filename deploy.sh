#!/bin/sh
echo "a" >> a.txt
kill -9 $(pgrep webserver)
cd ~/auto-deploy
git pull
cd webserver
./webserver &