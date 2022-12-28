init-db:
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
db-migrate:
	docker compose run --rm flyway migrate
db-migrate-info:
	docker compose run --rm flyway info
model-generate-docker:
	sqlboiler psql -c config/sqlboiler-docker.toml psql
model-generate-local:
	sqlboiler psql -c config/sqlboiler.toml psql
