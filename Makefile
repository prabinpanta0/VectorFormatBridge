# VectorFormatBridge Makefile

.PHONY: build clean test demo install

# Build the application
build:
	go build -o vectorformatbridge cmd/vectorformatbridge/main.go

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o dist/vectorformatbridge-linux-amd64 cmd/vectorformatbridge/main.go
	GOOS=windows GOARCH=amd64 go build -o dist/vectorformatbridge-windows-amd64.exe cmd/vectorformatbridge/main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/vectorformatbridge-darwin-amd64 cmd/vectorformatbridge/main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/vectorformatbridge-darwin-arm64 cmd/vectorformatbridge/main.go

# Clean build artifacts
clean:
	rm -f vectorformatbridge
	rm -rf dist/
	rm -f demo.svg demo.egf demo_converted.svg demo.egfb demo_decoded.egf

# Run tests
test:
	go test ./...

# Run demo
demo: build
	./vectorformatbridge demo

# Install to GOPATH/bin
install:
	go install cmd/vectorformatbridge/main.go

# Initialize module
init:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Create distribution directory
dist-dir:
	mkdir -p dist

# Show help
help:
	@echo "VectorFormatBridge Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  build-all  - Build for multiple platforms"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  demo       - Run demo"
	@echo "  install    - Install to GOPATH/bin"
	@echo "  init       - Initialize Go module"
	@echo "  fmt        - Format code"
	@echo "  lint       - Lint code (requires golangci-lint)"
	@echo "  help       - Show this help"
