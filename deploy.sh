#!/bin/bash
docker pull nazmialtun/testokur-light:latest
docker stop testokur-light && docker rm --force testokur-light
docker run -d  \
    --env-file  /home/env/testokur-light.env \
	--name testokur-light \
	--restart=always  \
	--network=testokur \
	--network-alias=testokur-light \
	nazmialtun/testokur-light:latest
echo Y | docker system prune
