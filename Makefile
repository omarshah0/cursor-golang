run: build
	./dist/main

build:
	mkdir -p dist
	go build -o dist/main .