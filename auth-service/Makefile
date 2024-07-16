SHELL := /bin/bash

include .env
export $(shell sed 's/=.*//' .env)

generate-env:
	@./generate_db_url.sh > .generated_env
	@echo "Generated .generated_env:"
	@cat .generated_env
	@source .generated_env

migrate-create: generate-env ### create new migration
	@source .generated_env && migrate create -ext sql -dir migrations 'centerx_db'
.PHONY: migrate-create

migrate-up: generate-env ### migration up
	@source .generated_env && migrate -path migrations -database "$$DB_URL" up
.PHONY: migrate-up

migrate-down: generate-env ### migration down
	@source .generated_env && migrate -path migrations -database "$$DB_URL" down
.PHONY: migrate-down

migrate-force: generate-env ### force migration to version 1
	@source .generated_env && migrate -path migrations -database "$$DB_URL" force 1
.PHONY: migrate-force

migrate-file: generate-env ### create a new migration file
	@source .generated_env && migrate create -ext sql -dir migrations/ -seq init_tables
.PHONY: migrate-file

proto:
	protoc --go_out=. \
  		--go-grpc_out=. \
  		--proto_path=protos \
  		auth.proto


give-permissions:
	- chmod +x additional.sh
	- chmod +x generate_db_url.sh

psql:
	psql -U postgres -d postgres