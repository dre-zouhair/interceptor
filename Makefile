.PHONY: format vet deps

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


run-dev:
	@echo "running protected services"
	@cd _dev/_cart-api && GOOS=linux go build -o ./build/protected-cart-api ./cmd/protected-api
	@docker compose -f ./_dev/protected.docker-compose.yml up -d --build
	@echo "running protected services is done"

clean-dev:
	@echo "cleaning protected services"
	@docker compose -f ./_dev/protected.docker-compose.yml down -v --rmi all
	@echo "cleaning services is done"