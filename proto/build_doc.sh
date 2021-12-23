#!/usr/bin/env bash

# change to project dir
CRTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $CRTDIR

echo "generating document..."
ret=0

fileList=""

for file in *.proto
do
    echo $file
    fileList="$fileList $file"
done

protoc -I=include -I=.  \
    --experimental_allow_proto3_optional \
    --openapiv2_out=. \
    --openapiv2_opt logtostderr=true,allow_merge=true,enums_as_ints=false,use_go_templates=true,merge_file_name=generated_doc \
    $fileList || ret=$?

mv generated_doc.swagger.json ../cmd/

exit $ret
