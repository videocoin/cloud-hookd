.NOTPARALLEL:

GOOS?=linux
GOARCH?=amd64

DOCKER_REGISTRY?=us.gcr.io/videocoin-network
APP_NAME?=hookd
VERSION?=$$(git describe --abbrev=0)-$$(git rev-parse --short HEAD)
IMAGE_TAG=${DOCKER_REGISTRY}/${APP_NAME}:${VERSION}

.PHONY: deploy

default: release

version:
	@echo ${VERSION}

image-tag:
	@echo ${IMAGE_TAG}

build:
	docker build -t ${IMAGE_TAG} -f Dockerfile .

build-bin:
	@echo "==> Building..."
	GOOS=${GOOS} GOARCH=${GOARCH} \
	go build -ldflags="-w -s -X main.Version=${VERSION}" -o bin/${APP_NAME} cmd/${APP_NAME}/main.go

test:
	@echo "No tests..."

push:
	@echo "==> Pushing ${APP_NAME} docker image..."
	docker push ${IMAGE_TAG}

release: build test push

