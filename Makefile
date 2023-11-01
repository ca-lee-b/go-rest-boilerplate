run:
	@go run cmd/main.go

build:
	@GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64 cmd/main.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 cmd/main.go

clean:
	@rm -r bin