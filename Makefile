


build:
	mkdir -p bin
	go build -o ./bin  ./...

check:
	go check ./...

tests:
	go tests ./...