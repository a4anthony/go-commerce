dev:
	air

swagger:
	swag init --dir ./,./handlers


# args: "--seed", "--refresh", "--seed --refresh"
db-migrate:
	go run ./migrate/migrate.go $(args)

db-migrate-seed:
	go run ./migrate/migrate.go --seed

db-migrate-refresh:
	go run ./migrate/migrate.go --refresh