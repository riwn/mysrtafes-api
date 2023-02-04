default: test

test: coverage.html

coverage.out:
	go test ./... -coverprofile=../cover/coverage.out

coverage.html: coverage.out
	go tool cover -html=../cover/coverage.out -o ../cover/coverage.html
	rm -f ../cover/coverage.out

run-api:
	go run ./cmd/api/main.go

build-api:
	go build -o ./dist/api cmd/api/main.go

fmt:
	go fmt ./...

tidy:
	go mod tidy
