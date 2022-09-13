#!/bin/bash

# https://earthly.dev/blog/docker-mysql/

docker run --name mysql -d \
   -p 3306:3306 \
   -e MYSQL_ROOT_PASSWORD=passwd \
   --restart unless-stopped \
   mysql:8

