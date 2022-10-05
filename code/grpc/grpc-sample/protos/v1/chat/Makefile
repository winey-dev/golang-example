# Makefile

all:
	protoc -I=. \
        -I=$(GOPATH)/src/grpc-example/protos \
        --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
        chat.proto;
clean:
	rm -rf *.go
