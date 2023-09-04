.PHONY: generate-proto
generate-proto:
	mkdir -p pkg
	protoc --go_out=./pkg --go_opt=paths=source_relative \
	--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
    api/imageproc_v1/imageprocapi.proto

.PHONY: clean-proto
clean-proto:
	rm -rf ./pkg

.PHONY: build
build:
	go build -o ./bin/ ./cmd/...

.PHONY: test
test:
	go test -v -race ./...







