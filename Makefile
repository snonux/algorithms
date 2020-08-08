all:
	go build
run:
	go run main.go
test:
	go test ./... -v
bench: sortbench
sortbench:
	go test -run=xxx -bench=. ./sort | tee sortbench.out
profile:
	go test -run=xxx -bench=BenchmarkQuickSort ./sort -memprofile memprofile.out -cpuprofile cpuprofile.out
	go tool pprof cpuprofile.out
