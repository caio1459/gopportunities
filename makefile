.PHONY: default run build test docs clean

# variaveis
APP_NAME=gopportunities

# tarefas
default: run

run: 
	@go run main.go
build:
	@go build
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs