.PHONY: format deps vet format-project deps-project vet-project .pidfile

format-project:
	@echo "formatting $(project) files..."
	@cd $(project) && go get golang.org/x/tools/cmd/goimports
	@cd $(project) && goimports -w -l .
	@cd $(project) && gofmt -s -w -l .
	@echo "formatting $(project) done!"

deps-project:
	@echo "installing $(project) dependencies..."
	@cd $(project) && go get ./...
	@echo "installing $(project) dependencies done"

vet-project:
	@echo "vetting $(project)..."
	@cd $(project) && go vet ./...
	@echo "vetting $(project) done"

test-project:
	@echo "starting $(project) tests..."
	@cd $(project) && go test ./... -coverprofile coverage.out
	@cd $(project) && go tool cover -html coverage.out -o coverage.html
	@echo "see coverage.out and coverage.html for $(project) code coverage"

win-build-project:
	@echo "building $(project)"
	@cd $(project) && go build -o ./build/app.exe ./cmd/api/main.go
	@echo "build $(project) done"

linux-build-project:
	@echo "building $(project)"
	@cd $(project) && go build -o ./build/app ./cmd/api/main.go
	@echo "build $(project) done"

run-project:
	@echo "building $(project)"
	@cd $(project) && go run ./cmd/api/main.go
	@echo "build $(project) done"

format:
	@make format-project project=_cart-api
	@make format-project project=_protection-api

deps:
	@make deps-project project=_cart-api
	@make deps-project project=_protection-api

vet:
	@make vet-project project=_cart-api
	@make vet-project project=_protection-api

tests:
	@make test-project project=_cart-api
	@make test-project project=_protection-api

win-build:
	@make win-build-project project=_cart-api
	@make win-build-project project=_protection-api

linux-build:
	@make linux-build-project project=_cart-api
	@make linux-build-project project=_protection-api

run-protection-api:
	@make run-project project=_protection-api

run-cart-api:
	@make run-project project=_cart-api

run-protected-cart-api:
	@cd _cart-api && go run ./cmd/protected-api/main.go

run-standalone-container:
	@echo "running standalone services"
	@cd _cart-api &&GOOS=linux go build -o ./build/app ./cmd/api/main.go
	@docker compose -f ./standalone.docker-compose.yml up -d --build
	@echo "standalone services are up"

clean-standalone-container:
	@echo "running standalone services"
	@docker compose -f ./standaone.docker-compose.yml down -v --rmi all
	@echo "standalone services are removed"

run-container:
	@echo "running standalone services"
	@docker compose -f docker-compose.yml up -d --build
	@echo "services are up"

clean-container:
	@echo "running standalone services"
	@docker compose -f docker-compose.yml down -v --rmi all
	@echo "services are removed"

run-protected-container:
	@echo "running protected services"
	@echo "building protected-cart-api"
	@cd _cart-api &&GOOS=linux go build -o ./build/protected-cart-api ./cmd/protected-api
	@echo "building protected-cart-api done"
	@docker compose -f ./protected.docker-compose.yml up -d --build
	@echo "protected services are up"

clean-protected-container:
	@echo "running protected services"
	@docker compose -f ./protected.docker-compose.yml down -v --rmi all
	@echo "protected services are removed"
