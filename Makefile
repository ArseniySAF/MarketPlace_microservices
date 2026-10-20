docker-up:
	docker compose up -d

auth-migrate-up:
	goose -dir ./gateway/migrations postgres "postgres://marketplace:marketplace_dev@localhost:5432/auth?sslmode=disable" up