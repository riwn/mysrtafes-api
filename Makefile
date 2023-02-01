default: test

first: api-build
	mkdir -p cover
	mkdir -p src/dist

test:
	docker compose run --rm gopher make

api: api-build
	docker compose up -d api


api-build:
	docker compose build api