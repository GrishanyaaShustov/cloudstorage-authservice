# Autoload .env
ifneq (,$(wildcard .env))
include .env
export $(shell sed 's/=.*//' .env)
endif

run:
	go run ./cmd/api

migrate:
	go run ./cmd/migrator

build:
	go build -o auth-service ./cmd/api