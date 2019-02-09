.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.DEFAULT_GOAL := docker

DOCKER_REGISTRY = us.gcr.io
CIRCLE_ARTIFACTS = ./bin
PROTOS_PATH = ./
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`
SERVICE_NAME = hookd
PROJECT_ID?=

VERSION=$$(git rev-parse --short HEAD)
IMAGE_TAG=$(DOCKER_REGISTRY)/${PROJECT_ID}/$(SERVICE_NAME):$(VERSION)
LATEST=$(DOCKER_REGISTRY)/${PROJECT_ID}/$(SERVICE_NAME):latest

IMAGE_TARBALL_PATH=$(CIRCLE_ARTIFACTS)/$(SERVICE_NAME)-$(VERSION).tar

main: build docker push

version:
	@echo $(VERSION)

image-tag:
	@echo $(IMAGE_TAG)

deps:
	 go mod vendor && go mod verify


build:
	cd cmd && xgo --targets=linux/amd64 -dest ../release -out $(SERVICE_NAME) .

build-alpine:
	export GOOS=linux
	export GOARCH=amd64
	export CGO_ENABLED=0
	go build -o bin/$(SERVICE_NAME) --ldflags '-w -linkmode external -extldflags "-static"' cmd/main.go

test:
	@echo "==> Running tests..."
	go test -v ./...

test-coverage:
	@echo "==> Running tests..."
	go test -cover ./...

docker:
	@echo "==> Docker building..."
	cd cmd && xgo -v --targets=linux/amd64 -dest ../release -out $(SERVICE_NAME) .
	docker build -t $(IMAGE_TAG) -t $(LATEST) . --squash
	docker push $(IMAGE_TAG)
	docker push $(LATEST)

proto-update:
	env GO111MODULE=on go get github.com/VideoCoin/common@latest
	env GO111MODULE=on go mod vendor
	env GO111MODULE=on go mod tidy

clean:
	rm -rf release/*


publish: package store clean
