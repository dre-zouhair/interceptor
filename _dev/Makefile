.PHONY: format deps vet format-project deps-project vet-project

format-project:
	@echo "formatting $(project) files..."
	@cd $(project) && go get golang.org/x/tools/cmd/goimports
	@cd $(project) && goimports -w -l .
	@cd $(project) && gofmt -s -w -l .
	@echo "formatting done!"

deps-project:
	@echo "installing $(project) dependencies..."
	@cd $(project) && go get ./...
	@echo "installing dependencies done"

vet-project:
	@echo "installing $(project) dependencies..."
	@cd $(project) && go go vet ./...
	@echo "installing dependencies done"

test-project:
	@echo "starting $(project) tests..."
	@cd $(project) && go test ./... -coverprofile coverage.out
	@cd $(project) && go tool cover -html coverage.out -o coverage.html
	@echo "see coverage.out and coverage.html for $(project) code coverage"

win-build-project:
	@echo "building $(project)"
	@cd $(project) && go build -o ./build/app.exe ./cmd/api/main.go
	@echo "build done"

linux-build-project:
	@echo "building $(project)"
	@cd $(project) && go build -o ./build/app ./cmd/api/main.go
	@echo "build done"

run-project:
	@echo "building $(project)"
	@cd $(project) &&  run ./... ./cmd/api/main.go
	@echo "build done"

format:
	@make format-project project=_cart-api
	@make format-project project=_protection-api

deps:
	@make deps-project project=_cart-api
	@make deps-project project=_protection-api

vet:
	@make vet-project project=_cart-api
	@make vet-project project=_protection-api

test:
	@make test-project project=_cart-api
	@make test-project project=_protection-api

win-build:
	@make win-build-project project=_cart-api
	@make win-build-project project=_protection-api

linux-build:
	@make linux-build-project project=_cart-api
	@make linux-build-project project=_protection-api

run:
	@make run-project project=_protection-api
	@make run-project project=_cart-api

run:
	@make run-project project=_protection-api
	@cd _cart-api && go run ./... ./cmd/protected-api/main.go

run-standalone-container:
	@echo "running standalone services"
	@docker compose -f docker-compose.standaone.yml up -d --build
	@echo "services are up"

clean-standalone-container:
	@echo "running standalone services"
	@docker compose -f docker-compose.standaone.yml down -v --rmi all
	@echo "services are up"

run-container:
	@echo "running standalone services"
	@docker compose -f docker-compose.yml up -d --build
	@echo "services are up"

clean-container:
	@echo "running standalone services"
	@docker compose -f docker-compose.yml down -v --rmi all
	@echo "services are up"



