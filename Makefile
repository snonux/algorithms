all:
	go build
run:
	go run main.go
test:
	go test ./... -v
bench:
	go test -run=xxx -bench=. ./... | tee bench.out
sortbench:
	go test -run=xxx -bench=. ./sort | tee sortbench.out
queuebench:
	go test -run=xxx -bench=. ./queue | tee queuebench.out
searchbench:
	go test -run=xxx -bench=. ./search | tee searchbench.out
profile:
	go test -run=xxx -bench=BenchmarkQuickSort ./sort -memprofile memprofile.out -cpuprofile cpuprofile.out
	go tool pprof cpuprofile.out
