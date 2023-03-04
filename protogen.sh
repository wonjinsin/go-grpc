mkdir -p dependency/proto-gen-gorm
git clone https://github.com/infobloxopen/protoc-gen-gorm.git proto/dependency/proto-gen-gorm

mkdir -p dependency/proto-gen-gorm/proto-gen
cd proto/dependency/proto-gen-gorm
make build-user-local


mkdir -p generate
protoc -I. -I./proto/dependency/proto/options --gorm_out="engine=postgres,enums=string,gateway:./generate" --go_out=./generate proto/model/*.proto
protoc -I. --go_out=./generate --go-grpc_out=./generate proto/controller/v1/*.proto

mv ./generate/phantom/model/* ./model
mv ./generate/phantom/controller/v1/* ./controller/v1
rm -rf ./generate
