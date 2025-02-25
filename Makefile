# Define binary name
BINARY_NAME = go-pgfmt
INSTALL_PATH = /usr/local/bin

# Build the binary
build:
	go build -o $(BINARY_NAME) .

# Install the binary globally
install: build
	mv $(BINARY_NAME) $(INSTALL_PATH)/
	chmod +x $(INSTALL_PATH)/$(BINARY_NAME)
	echo "Installation complete! Run 'go-pgfmt --help' to verify."

# Clean up generated files
clean:
	rm -f $(INSTALL_PATH)/$(BINARY_NAME)

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...
