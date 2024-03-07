# Define the port to be used (default is 8080)
PORT := :8080

run:
	templ generate
	PORT=$(PORT) 
	go run cmd/main.go

fmt:
	go fmt cmd/main.go
	go fmt handlers/**
	go fmt services/**