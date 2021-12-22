#!/bin/bash

# change to proto dir
CRTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $CRTDIR

#init
DST_DIR="generated_go"

test -e "generated_go" && rm -rf $DST_DIR

mkdir -p $DST_DIR

for file in *.proto
do
  echo $file
  protoc -I=include -I=. \
    --go_out=$DST_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$DST_DIR --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$DST_DIR --grpc-gateway_opt=logtostderr=true,paths=source_relative \
    `find "${file}" -maxdepth 1 -name '*.proto'`
done