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
        user.proto;
swag-all:
	protoc -I=. \
        -I=$(GOPATH)/src/grpc-example/protos \
        --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
		--swagger_out=logtostderr=true:. \
		--openapiv2_out ./gen/openapiv2 \
		--openapiv2_opt logtostderr=true \
        user.proto;
clean:
	rm -rf *.go *.json
