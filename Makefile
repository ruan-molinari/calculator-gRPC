generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run server/main.go

run_grpcui:
	grpcui --plaintext localhost:8000
