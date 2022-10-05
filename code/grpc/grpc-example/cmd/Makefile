all:
	rm -rf grpc-server
	go build -o server

install:
	cp -rp grpc-server ../build/docker/grpc-server
clean:
	rm -rf grpc-server cmd server
