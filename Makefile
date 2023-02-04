default: test

first: api-build
	mkdir -p cover
	mkdir -p src/dist
	cp compose.override.yml.sample compose.override.yml

test:
	docker compose run --rm gopher make

api: api-build
	docker compose up -d api

api-build:
	docker compose build api

down:
	docker compose down
