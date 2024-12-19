DB_USER := postgres
DB_PASSWORD := postgres
DB_NAME := notification
DB_HOST := localhost
DB_PORT := 5432
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

deploy:
	@docker-compose -f ./docker-compose.yml up

undeploy:
	@docker-compose -f ./docker-compose.yml down