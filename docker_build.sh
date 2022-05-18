#!/bin/bash

version=$1

docker build -t jw-base:$version .

# shellcheck disable=SC2086
docker tag jw-base:$version www.jwdouble.top:10443/k8s/jw-base:$version

# shellcheck disable=SC2086
docker push www.jwdouble.top:10443/k8s/jw-base:$version