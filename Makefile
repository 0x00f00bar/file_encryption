run:
	@go run main.go

build: clean
	@echo "Building for release"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/linux-amd64 main.go
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/darwin-amd64 main.go
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o ./bin/darwin-arm64 main.go
	@GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/win-amd64.exe main.go
	@tar -czf ./bin/file_encryption-linux-amd64.tar ./bin/linux-amd64 && rm ./bin/linux-amd64
	@tar -czf ./bin/file_encryption-darwin-amd64.tar ./bin/darwin-amd64 && rm ./bin/darwin-amd64
	@tar -czf ./bin/file_encryption-darwin-arm64.tar ./bin/darwin-arm64 && rm ./bin/darwin-arm64
	@tar -czf ./bin/file_encryption-win-amd64.tar ./bin/win-amd64.exe && rm ./bin/win-amd64.exe
	@echo "Done"

clean:
	@echo "Cleaning..."
	@rm ./bin/* || echo "Empty directory"
