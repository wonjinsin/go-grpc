## 참고
<https://devjin-blog.com/golang-grpc-server-1/>


## 명령어
```
go mod init go-grpc
go get -u google.golang.org/grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative protos/v1/user/user.proto
```