.PHONY: test all clean

all:
	go build -o new ./cmd/new/main.go

clean:
	rm new

test:
	go test ./...
