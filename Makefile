
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/*/*.proto

