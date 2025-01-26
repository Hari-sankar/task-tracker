include .env

.PHONY: test-env migrate-up migrate-down migrate-create

# Command to verify env variables are loaded
test-env:
	@echo "Database URL: $(DB_SOURCE)"


migrate-up:
	migrate -path db/migrations -database "$(DB_SOURCE)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_SOURCE)" down

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-force:
	migrate -path db/migrations -database "$(DB_SOURCE)" force $(version)

migrate-version:
	migrate -path db/migrations -database "$(DB_SOURCE)" version

migrate-down-steps:
	migrate -path db/migrations -database "$(DB_SOURCE)" down $(steps)
