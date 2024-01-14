BINARY_NAME=managment_system

build:
	go build -o ./bin/$(BINARY_NAME) ./cmd/app

run: build
	./bin/$(BINARY_NAME)

test:
	go test ./...

clean:
	rm -f ./bin/$(BINARY_NAME)

start:
	docker-compose up -d --build

stop:
	docker-compose down