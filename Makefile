build-darwin-amd64:
	mkdir -p build/darwin/amd64 && env GOOS=darwin GOARCH=amd64 go build -o build/darwin/amd64/restcups server.go
build-darwin-arm64:
	mkdir -p build/darwin/arm64 && env GOOS=darwin GOARCH=arm64 go build -o build/darwin/arm64/restcups server.go
build-linux-amd64:
	mkdir -p build/linux/amd64 && env GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/restcups server.go
build-linux-arm64:
	mkdir -p build/linux/arm64 && env GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/restcups server.go
