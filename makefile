run-test:
	go test ./... -covermode=count -coverprofile ./coverage.out

run-fmt:
	go fmt

check-swagger:
	which swagger || (GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=on swagger generate spec -o ./docs/swagger.yaml --scan-models

setting-prepare:
	go mod tidy
	go mod vendor

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build .

install-app: build-app
	go install