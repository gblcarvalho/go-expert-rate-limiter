start:
	docker compose --env-file ./.env up -d

stop:
	docker compose --env-file ./.env down

test:
	./test.sh
