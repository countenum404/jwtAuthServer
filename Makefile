app = jwtAuthServer

up:
	docker-compose up --build $(app)

clean:
	docker compose rm -f -s



