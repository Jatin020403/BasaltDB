.PHONY: build run test

build: 
	@go build -o bin/basaltdb

run: build
	@./bin/basaltdb $(filter-out $@,$(MAKECMDGOALS))

test:
	@go test ./... 

%:
	@: