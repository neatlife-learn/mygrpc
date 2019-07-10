#!/bin/bash

for file in ./protos/*.proto; do
    echo "生成对应的php类文件: $file"
    protoc -I ./protos --php_out=./php/pbf --grpc_out=./php/pbf --plugin=protoc-gen-grpc=$(which grpc_php_plugin) $file
done

protoc -I ./protos --go_out=plugins=grpc:./go/pbf ./protos/*.proto
