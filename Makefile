compose-up:
	docker-compose -f local-dev/docker-compose.yaml --profile infra up -d

compose-down:
	docker-compose -f local-dev/docker-compose.yaml --profile infra down