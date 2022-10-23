PWD = $(shell pwd)
ALL_DOCKER = $(shell docker ps -q -a)

.PHONY: build all clean run stop test

build: test
	docker build . -t signing-service

run: build
	docker run -d -v ${PWD}/target:/data -p 3000:3000 signing-service > PID

stop:
	docker stop `cat PID`

all: build run

request:
	curl -X POST http://localhost:3000/hello -d '{"name": "world"}'
test:
	go test -v ./...

clean:
	docker rm  ${ALL_DOCKER}
	rm -rf target
	rm PID