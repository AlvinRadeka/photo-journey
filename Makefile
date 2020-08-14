.PHONY: build
build:
	go build -v -o app *.go

.PHONY: run
run:
	@go build -v -o app *.go
	@./app
