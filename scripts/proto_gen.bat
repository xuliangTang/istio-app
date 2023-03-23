protoc320 --proto_path=protos/userProto --go_out=./ Models.proto
protoc320 --proto_path=protos/userProto --go_out=./ UserKind.proto
protoc320 --proto_path=protos/userProto --go-grpc_out=./ UserSvc.proto

protoc320 --proto_path=protos/prodProto --go_out=./ ProdModels.proto
protoc-go-inject-tag -input="./internal/prodApp/v1/*.pb.go"

protoc320 --proto_path=protos/orderProto --go_out=./ OrderModels.proto
protoc-go-inject-tag -input="./internal/orderApp/v1/*.pb.go"