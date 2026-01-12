include .envrc

MIGRATIONS_PATH= /Users/ayomideademilua/Development/go_crash_course/go_social/cmd/migrate/migrations


.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))


.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=disable up


.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=disable down $(filter-out $@,$(MAKECMDGOALS))


.PHONY: seed
seed:
	@direnv allow /Users/ayomideademilua/Development/go_crash_course/go_social/.envrc
	@go run cmd/migrate/seed/main.go