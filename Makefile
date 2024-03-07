run:
	templ generate
	go run cmd/main.go

fmt:
	go fmt cmd/main.go
	go fmt handlers/**
	go fmt services/**