goos: darwin
goarch: amd64
pkg: github.com/knight42/benchmarks/concat_slices
BenchmarkConcatCopyPreAllocate-4   	30000000	        50.5 ns/op	      64 B/op	       1 allocs/op
BenchmarkConcatAppend-4            	20000000	       106 ns/op	     112 B/op	       3 allocs/op
PASS
ok  	github.com/knight42/benchmarks/concat_slices	3.824s
