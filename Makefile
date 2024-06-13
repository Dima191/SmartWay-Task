migrate_down:
	migrate --path ./migrations --database {DB_URL} down
migrate_up:
	migrate --path ./migrations --database {DB_URL} up

run:
	go build ./cmd/app
	./app

.PHONY: migrate_up, migrate_down, run

.DEFAULT_GOAL = run