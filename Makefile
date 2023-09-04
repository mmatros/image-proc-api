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

.PHONY: generate-mock
generate-mock:
	mockgen -source ./pkg/api/imageproc_v1/imageprocapi_grpc.pb.go \
	-destination ./pkg/api/imageproc_v1/mocks/imageprocapi_grpc.pb.go \
	-package imageprocapi_v1


.PHONY: test
test:
	go test -v -race ./...







