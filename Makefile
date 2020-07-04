.PHONY: test all clean

all:
	# go build -o new ./cmd/new/main.go
	pyinstaller ./src/new.py

clean:
	# rm new
	rm -rf ./build
	rm -rf ./dist

test:
	# go test ./...
	pytest
