protoc320 --proto_path=protos/userProto --go_out=./ Models.proto
protoc320 --proto_path=protos/userProto --go_out=./ UserKind.proto
protoc320 --proto_path=protos/userProto --go-grpc_out=./ UserSvc.proto

protoc320 --proto_path=protos/orderProto --go_out=./ Models.proto
protoc-go-inject-tag -input="./internal/orderApp/v1/*.pb.go"
