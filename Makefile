.PHONY: all api clean build_server build_client

api/api.pb.go: api/api.proto
	protoc -I api/ \
    		-I${GOPATH}/src \
    		--go_out=plugins=grpc:api \
    		api/api.proto

all: build_client build_server

dep:
	go get -v -d ./...

build_server: dep api
	go build -i -v -o bin/server github.com/benjaminhadfield/grpc/server

build_client: dep api
	go build -i -v -o bin/client github.com/benjaminhadfield/grpc/client

api: api/api.pb.go

clean:
	rm api/api.pb.go
