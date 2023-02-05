## 참고
<https://devjin-blog.com/golang-grpc-server-1/>


## 명령어
```
go mod init go-grpc
go get -u google.golang.org/grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative protos/v1/user/user.proto
```

## 테스트 (ezy 이용)
- ezy collection에 user.proto와 post.proto 추가

## go grpc for json
```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```