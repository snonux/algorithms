all:
	go build
run:
	go run main.go
test:
	go test ./... -v
bench: sortbench queuebench
sortbench:
	go test -run=xxx -bench=. ./sort | tee sortbench.out
queuebench:
	go test -run=xxx -bench=. ./queue | tee queuebench.out
	go test -run=xxx -bench=. ./queue | tee queuebench.out
profile:
	go test -run=xxx -bench=BenchmarkQuickSort ./sort -memprofile memprofile.out -cpuprofile cpuprofile.out
	go tool pprof cpuprofile.out
