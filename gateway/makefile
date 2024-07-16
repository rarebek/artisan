generate-rpc:
	protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/auth-protos/auth.proto \
	protos/product-protos/product.proto


swag:
	swag init -g api/api.go 