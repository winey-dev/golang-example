# 1. proto 컴파일러 설치하기
> https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.1

공식 릴리즈에서 각 OS맞는 파일을 다운로드 

압축 해제 후 PATH=$PROTOC_PATH/bin 설정  
# 2. Go Code Gen 설치
~~~
> $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
> $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
~~~


# 3. protoc 명령어를 사용하여 proto 컴파일 하기 
~~~
protoc --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       user.proto

~~~
~~~
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       user.proto

~~~

# 4. 공식 가이드 문서
 https://developers.google.com/protocol-buffers/docs/overview
