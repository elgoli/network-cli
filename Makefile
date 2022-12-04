.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: build
build:
	GOOS=linux go build -o bin/network-cli -v main.go