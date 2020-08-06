all:
	go build
run:
	go run main.go
test:
	go test ./... -v
bench:
	go test -run=10 -bench=. ./... -v
profile:
	go test -bench=QuickSort ./sort -memprofile memprofile.out -cpuprofile profile.out
	go tool pprof profile.out
