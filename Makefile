# Define the port to be used (default is 8080)
PORT := :8080

run:
	templ generate
	PORT=$(PORT) go run cmd/main.go

format:
	go fmt views/**