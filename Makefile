.PHONY: build run test

build: 
	@go build -o bin/basaltdb
	@./bin/basaltdb $(filter-out $@,$(MAKECMDGOALS))

run: build
	@./bin/basaltdb $(filter-out $@,$(MAKECMDGOALS))

test:
	@go test ./... 

%:
	@: