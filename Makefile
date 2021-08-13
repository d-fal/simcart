
GOPATH:=$(shell go env GOPATH)

proto:
	mkdir -p api/swagger


	git submodule update --recursive --remote
	
	rm -rf api/pb api/swagger
	mkdir -p api/pb api/swagger
	
	protoc -I pb/aux \
	-I pb/proto \
	--go_out api/pb \
	--openapiv2_out=logtostderr=true,repeated_path_param_separator=ssv:./api/swagger \
	--openapiv2_opt use_go_templates=true \
	--openapiv2_opt logtostderr=true \
	--openapiv2_opt use_go_templates=true \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--go_opt paths=source_relative \
	--go-grpc_out api/pb \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out api/pb \
	--grpc-gateway_opt paths=source_relative \
	pb/proto/**/*.proto \
	pb/proto/**/**/*.proto

.PHONY: build
build:
	go build -o micro *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: vendor
vendor:
	go get ./...
	go mod vendor
	go mod verify


config:
	cp -rf ./config.example.yaml ./config.yaml

.PHONY: config
.PHONY: proto