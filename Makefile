.PHONY: build run test

build: 
	@go build -o bin/basaltdb

run: build
	@./bin/basaltdb $(filter-out $@,$(MAKECMDGOALS))

test:
	@go test ./... 

insert_data:
	@for i in $(shell seq 1 100); do \
		./bin/basaltdb insert key$$i value$$i; \
	done

%:
	@:

