# Makefile to build, run and clean the grpc-server project

.PHONY: build run clean

build:
	docker build -t webserver .

run:
	docker run -p 50051:50051 webserver

clean:
	docker rmi webserver
