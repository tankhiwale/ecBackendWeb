b:
	@go build -o ./bin/web ./...

run: b
	@ ./bin/web
