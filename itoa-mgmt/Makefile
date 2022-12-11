clean:
	rm pb/*
	rm swagger/*

gen:
	 source $HOME/.bashrc && protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger