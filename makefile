#GOBIN?=.
GOBIN?=/src/bin
GO111MODULE?=on
PROJECT_NAME?=cosmoappy

run-test:
	go clean -testcache
	go test ./... -covermode=count -coverprofile ./coverage.out

run-fmt:
	go fmt

check-swagger: setting-prepare
	which swagger || (GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	swagger generate spec -o ./swagger.yaml --scan-models

yamlfmt:
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	$(HOME)/go/bin/yamlfmt

setting-prepare:
	go mod tidy
	go mod vendor

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build -mod=vendor -o $(GOBIN)/$(1)

install-app: build-app
	go install

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start: docker-build
	docker run -v $(PWD)/test:/etc/$(1)/conf -p 8080:8080 $(PROJECT_NAME):latest -conf=/etc/$(1)/conf/conf_docker.yaml
