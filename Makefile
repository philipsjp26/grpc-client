#!/bin/bash
run: 
	@go run main.go

gen-proto:
	@protoc --proto_path=. *.proto --go_out=common --go-grpc_out=common