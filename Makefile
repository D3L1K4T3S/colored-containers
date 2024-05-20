PATH = ./src/

build:
	go build -o $(PATH)/bin/main main.go

run:
	go run main.go

test:
	go test $(PATH)

clean:
	rm -rf $(PATH)/bin

.PHONY: build

