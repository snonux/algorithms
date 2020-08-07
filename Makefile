all:
	go build
run:
	go run main.go
test:
	go test ./... -v
bench:
	go test -run=xxx -bench=. ./... -v
profile:
	go test -run=xxx -bench=Quick2Sort ./sort -memprofile memprofile.out -cpuprofile cpuprofile.out
	go tool pprof cpuprofile.out
