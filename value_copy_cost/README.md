```
goos: darwin
goarch: amd64
pkg: github.com/knight42/benchmarks/value_copy_cost
Benchmark_Loop-4               	 1000000	      1511 ns/op
Benchmark_Range_OneIterVar-4   	 1000000	      1509 ns/op
Benchmark_Range_TwoIterVar-4   	  500000	      2816 ns/op
PASS
ok  	github.com/knight42/benchmarks/value_copy_cost	4.499s
```
