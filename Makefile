build:
	docker compose build

run:
	docker compose up

stop:
	docker compose down

gen:
	protoc -I=proto --go_out=proto/pb --go-grpc_out=proto/pb proto/playlist.proto

evans:
	evans proto/playlist.proto

.PHONY: build run stop gen evans 