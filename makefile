#GOBIN?=.
GOBIN?=/src/bin
GO111MODULE?=on
PROJECT_NAME?=cosmoappy
PROJECT_PATH?=CosmoAppy

setting-prepare:
	go mod tidy
	go mod vendor

install-devtools: setting-prepare
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	go install github.com/segmentio/golines@latest

install-swagger: setting-prepare
	which swagger || (GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger)

run-test:
	go clean -testcache
	go test ./... -covermode=count -coverprofile ./coverage.out

run-fmt: install-devtools
	go fmt
	$(HOME)/go/bin/yamlfmt
	$(HOME)/go/bin/golines . -w

swagger: install-swagger
	swagger generate spec -o test/swagger.yaml --scan-models

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build -mod=vendor -o $(GOBIN)/$(PROJECT_PATH)

install-app: build-app
	go install

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker run \
		--name $(PROJECT_NAME) \
		-v $(PWD)/test:/etc/$(PROJECT_PATH)/conf \
		-e PASS_DB=${PASS_DB} -e TOKEN_API_TEST=${TOKEN_API_TEST} \
		-p 8080:8080 -p 8181:8181 \
		$(PROJECT_NAME):latest \
		-conf=/etc/$(PROJECT_PATH)/conf/conf_docker.yaml \
		-swagger=/swagger.yaml
