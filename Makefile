
format:
	@echo "formatting files..."
	@go get golang.org/x/tools/cmd/goimports
	@goimports -w -l .
	@gofmt -s -w -l .
	@echo "formatting done!"

vet:
	@echo "vetting..."
	@go vet ./...
	@echo "vetting done!"

deps:
	@echo "installing dependencies..."
	@go get ./...