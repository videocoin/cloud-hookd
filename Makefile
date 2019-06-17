.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.DEFAULT_GOAL := docker

DOCKER_REGISTRY = us.gcr.io
CIRCLE_ARTIFACTS = ./bin
SERVICE_NAME = hookd

PROJECT_ID= videocoin-network
VERSION=$$(git describe --abbrev=0)-$$(git rev-parse --short HEAD)
IMAGE_TAG=$(DOCKER_REGISTRY)/$(PROJECT_ID)/$(SERVICE_NAME):$(VERSION)
LATEST=$(DOCKER_REGISTRY)/$(PROJECT_ID)/$(SERVICE_NAME):latest

main: build docker push

version:
	@echo $(VERSION)

image-tag:
	@echo $(IMAGE_TAG)

deps:
	env GO111MODULE=on go mod vendor

build:
	export GOOS=linux
	export GOARCH=amd64
	export CGO_ENABLED=0
	@echo "==> Building..."
	@go build -a -installsuffix cgo -ldflags="-w -s" -o bin/$(SERVICE_NAME) cmd/main.go


test:
	@echo "==> Running tests..."
	@go test -v ./...

test-coverage:
	@echo "==> Running tests..."
	go test -cover ./...

docker:
	@echo "==> Docker building..."
	@docker build -t $(IMAGE_TAG) -t $(LATEST) .

push:
	@docker push $(IMAGE_TAG)
	@docker push $(LATEST)

clean:
	rm -rf release/*


publish: package store clean
