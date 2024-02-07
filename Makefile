.PHONY: format vet deps run-dev run-web-app run-protection-api run-protected-cart-api clean-dev

format:
	@echo "formatting files..."
	@go get golang.org/x/tools/cmd/goimports
	@goimports -w -l .
	@gofmt -s -w -l .
	@cd _dev && make format
	@echo "formatting done!"

vet:
	@echo "vetting..."
	@go vet ./...
	@cd _dev && make vet
	@echo "vetting done!"

deps:
	@echo "installing dependencies..."
	@go get ./...
	@cd _dev && make deps
	@echo "installing dependencies done!"

run-protection-api:
	@cd _dev && make run-protection-api

run-protected-cart-api:
	@cd _dev && make run-protected-cart-api

run-web-app:
	@cd _dev && make run-web-app

run-dev:
	@echo "running protected services"
	@cd _dev && make deps
	@cd _dev/_cart-api && GOOS=linux go build -o ./build/protected-cart-api ./cmd/protected-api
	@docker compose -f ./_dev/protected.docker-compose.yml up -d --build
	@echo "running protected services is done"

clean-dev:
	@echo "cleaning protected services"
	@docker compose -f ./_dev/protected.docker-compose.yml down -v --rmi all
	@echo "cleaning services is done1"