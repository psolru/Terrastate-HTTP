#!/bin/bash

# Todo: Since i recently learned how to do it, make a docker-compose.yml out of it :D

echo "Kill existing Docker container"
docker stop terrastate-http \
  && docker rm terrastate-http

echo "Spin up new Docker container"
time docker build . -t terrastate-http \
  && docker run --detach \
    --name terrastate-http \
    --publish 8080:8080 \
    --volume $(pwd)/examples/configs/config.json:/mnt/configs/config.json \
    --volume $(pwd)/examples/sqlite3/:/mnt/sqlite3/ \
    terrastate-http

docker logs -f terrastate-http
