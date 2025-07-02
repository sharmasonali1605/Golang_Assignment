# Golang_Assignment

To run server use this command , server is running at port 8072
# go run server/*.go 

To run clinet use command
# go run client/main.go

To run test use this command
# go test ./service

Also I have captured a vedio demostrating response of APIs through Postman whihc is shared in a mail.

# comnad to generate protofiles.
protoc --go_out=. --go-grpc_out=. --proto_path=blogpb blogpb/blog.proto

#  protoc --go_out=. --go-grpc_out=. --proto_path=/Users/sonali2302/Golang_Assignment/blogpb/blog.proto