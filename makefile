GO111MODULE?=on
PROJECT_NAME?=cosmoappy
PROJECT_PATH?=CosmoAppy

setting-prepare:
	go mod tidy
	go mod vendor

install-devtools: setting-prepare
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	go install github.com/segmentio/golines@latest
	which swagger || (GO111MODULE=on go install github.com/go-swagger/go-swagger/cmd/swagger@latest)

run-test:
	go clean -testcache
	go test ./... -covermode=count -coverprofile ./coverage.out

swagger:
	$(HOME)/go/bin/swagger generate spec -o ./conf/swagger.yaml --scan-models

run-fmt:
	go fmt
	$(HOME)/go/bin/yamlfmt conf/
	$(HOME)/go/bin/golines . -w

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build -mod=vendor

install-app: build-app
	go install

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker run --rm \
		--name $(PROJECT_NAME) \
		-v $(PWD)/conf:/etc/$(PROJECT_PATH)/conf \
		-e PASS_DB=${PASS_DB} -e TOKEN_API_TEST=${TOKEN_API_TEST} \
		-p 8080:8080 \
		$(PROJECT_NAME):latest \
		-conf=/etc/$(PROJECT_PATH)/conf/conf_docker.yaml
