#!/bin/bash

REPOSITORY=${1}

aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 504784160859.dkr.ecr.ap-northeast-1.amazonaws.com
docker build -t sample .
docker tag sample:latest 504784160859.dkr.ecr.ap-northeast-1.amazonaws.com/${REPOSITORY}:latest
docker push 504784160859.dkr.ecr.ap-northeast-1.amazonaws.com/${REPOSITORY}:latest
