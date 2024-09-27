app = application

up:
	docker-compose up --build $(app)

clean:
	docker compose rm -f -s

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:1234@localhost/jwt?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:1234@localhost/jwt?sslmode=disable' down
