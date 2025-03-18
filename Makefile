include .env

run:
	go run cmd/web/main.go

start-nodemon:
	nodemon --exec go run cmd/web/main.go --signal SIGTERM

test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out && rm -f coverage.out

create-mock:
	mockery --all --case underscore --output ./mocks