# Makefile

all:
	protoc -I=. \
        -I=$(GOPATH)/src/grpc-example/protos \
        --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
        --grpc-gateway_out . \
        --grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		v2chat.proto;
clean:
	rm -rf *.go
