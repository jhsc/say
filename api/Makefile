build:
	docker run --rm \
    -u $(id -u):$(id -g) \
    -v $(PWD):/src:rw \
    -w /src \
    gnhuy91/protobuf-go \
    --go_out=plugins=grpc:. *.proto