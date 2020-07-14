all:
	go build
run:
	go run main.go
test:
	go test ./... -v
bench:
	go test -run=10 -bench=. ./... -v
