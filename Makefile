.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.DEFAULT_GOAL := main

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
	@echo "==> Building..."
	export GOOS=linux
	export GOARCH=amd64
	export CGO_ENABLED=0
	go build -o bin/$(SERVICE_NAME) cmd/main.go

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
	docker build  -t ${IMAGE_TAG} -t $(LATEST) . --squash

push:
	@echo "==> Pushing $(IMAGE_TAG)..."
	gcloud docker -- push $(IMAGE_TAG)
	gcloud docker -- push $(LATEST)


docker-save:
	@echo "==> Saving docker image tarball..."
	gcloud auth configure-docker --quiet
	docker save -o $(IMAGE_TARBALL_PATH) $(IMAGE_TAG)

proto-update:
	env GO111MODULE=on go get github.com/VideoCoin/common@latest
	env GO111MODULE=on go mod vendor
	env GO111MODULE=on go mod tidy

package:
	cd cmd && xgo --targets=linux/amd64 -dest ../release -out $(SERVICE_NAME) .
	cp keys/$(SERVICE_NAME).key release
	tar -C release -cvjf release/$(VERSION)_$(SERVICE_NAME)_linux_amd64.tar.bz2 $(SERVICE_NAME)-linux-amd64 $(SERVICE_NAME).key

store:
	gsutil -m cp release/$(VERSION)_$(SERVICE_NAME)_linux_amd64.tar.bz2 gs://${RELEASE_BUCKET}/$(SERVICE_NAME)/
	gsutil -m cp gs://${RELEASE_BUCKET}/$(SERVICE_NAME)/$(VERSION)_$(SERVICE_NAME)_linux_amd64.tar.bz2 gs://${RELEASE_BUCKET}/$(SERVICE_NAME)/latest_$(SERVICE_NAME)_linux_amd64.tar.bz2

clean:
	rm -rf release/*


publish: package store clean
