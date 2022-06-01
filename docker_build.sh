#!/bin/bash

make

version=$(date "+%Y%m%d%H%M")

docker build -t jw-base:$version .

# shellcheck disable=SC2086
docker tag jw-base:$version www.jwdouble.top:10443/k8s/jw-base:$version

# shellcheck disable=SC2086
docker push www.jwdouble.top:10443/k8s/jw-base:$version

sed -i "s/jw-base:.*/jw-base:${version}/g" k8s.yaml

scp ./k8s.yaml root@150.158.7.96:/root/app/k8s/dep/jw/base/