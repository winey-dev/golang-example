# Learning Objectives
===================
- [x] Golang으로 gRPC Server 를 구현해본다
- [x] proto-gateway를 통해 Restapi와 연동하기
- [x] proto-middleware 사용 및 구현 해보기

# grpc-example 
## protoc 설치 
~~~bash
:> git clone https://github.com/protocolbuffers/protobuf.git  
:> cd protobuf
:> ./configure
:> sudo make 
:> sudo make check
:> sudo make install
~~~

## Golang Protoc gen 설치 
~~~bash
:> go install google.golang.org/protobuf/cmd/protoc-gen-go \
   google.golang.org/grpc/cmd/protoc-gen-go-grpc
~~~

# proto-gateway

## Gateway Protobuffer 사용을 위한 PKG Download
~~~bash
:> go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
          github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
~~~

## Gateway proto buffer 컴파일 오류 해결 내용
### 1) not found google/api/annotaions.proto 에러 메시지 해결 방법
~~~bash
# 현재 경로가 $GOPATH/src/grpc-example 이라고 가정 할 경우
:> git clone https://github.com/googleapis/googleapis.git
:> cp googleapis/google/api/annotaions.proto  ${GOPATH}/src/grpc-example/protos/google/api/annotaions.proto
:> cp googleapis/google/api/http.proto ${GOPATH}/src/grpc-example/google/api/http.proto
~~~

### 2) not found google/protobuf/descriptor.proto 에러 메시지 해결 방법 
   
#### local dir 확인 방법
~~~bash
:> cd /usr/local/include/google #해당 파일의 권한 변경 해주기
:> chmod -R 755 /usr/local/include/google
~~~

#### /usr/local/include/google 이 없을 경우 
~~~bash
:> git clone https://github.com/protocolbuffers/protobuf.git  
:> cp -rp protobuf/src/google /usr/local/include/google
~~~

# proto-middleware
grpc-ecosystem/go-grpc-middleware 에서 제공해주는 middleware 사용해보기

custerm middleware 구현 하기

# protos version 정보
|경로|grpc-server|grpc-gateway|stream|
|:---:|:---:|:---:|:---:|
|protos/v1/user|o|x|x|
|protos/v1/chat|o|x|o|
|protos/v2/user |o|o|x|

# project 경로 설명
|Path                           |Description                     |
|:------------------------------|:-------------------------------|
|cmd/main.go                    |프로젝트 메인 코드              |
|cmd/grpc/server.go             |grpc server 구현 부분           |
|cmd/grpc/gateway.go            |grpc gateway 구현 부분          |
|cmd/grpc/middleware.go         |grpc middleware 구현 부분       |
|internal/$(version)/$(usercase)|proto에 정의한 Serivce 구현 부분|
|protos/$(version)/$(usercase)  |proto buffer 정의               |

* * *
# 설치 및 테스트 방법
## Install 
~~~bash
:> cd $GOPATH/src
:> git clone https://github.com/yiaw/grpc-example.git
:> cd grpc-example/cmd
:> make
:> ./grpc-server
2022/01/07 14:34:23 start gRPC Server on 8090 port, enableTLS=false
2022/01/07 14:34:23 HTTP Server GRPC Gateway on http://0.0.0.0:8080
~~~


* * *
# gRPC-GO 를 사용하면서 궁금한점
1. Golang에서는 Interface{} 변수를 사용 할 수 있는데 Protobuf와 매칭되는 특정 예약어가 있는가?
   proto syntax3 에서 [Any](https://developers.google.com/protocol-buffers/docs/proto3#any)를 지원 한다.   
      
2. gRPC에서 HTTP Header의 값을 얻어오는 방법
   Context 객체에서 얻어 올 수 있다. [링크](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)

3. Protobuf 정의 시 GET, DELETE에 Body를 정의 하는 방법
   GET은 공식적으로 지원하지 않는다.   
   DELETE의 경우 Proto를 컴파일 시 `--grcp-gateway_out allow_delete_body=true:.`를 추가 한다.
