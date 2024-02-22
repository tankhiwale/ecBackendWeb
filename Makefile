b:
	@go fmt ./...
	@go build -o ./bin/web ./cmd/main.go

run: b
	@ ./bin/web
