```
goos: darwin
goarch: amd64
pkg: github.com/knight42/benchmarks/concat_slices
BenchmarkConcatCopyPreAllocate-4   	20000000	        72.1 ns/op	     176 B/op	       1 allocs/op
BenchmarkConcatAppend-4            	10000000	       154 ns/op	     288 B/op	       4 allocs/op
BenchmarkConcatWriteBuffer-4       	10000000	       226 ns/op	     352 B/op	       2 allocs/op
PASS
ok  	github.com/knight42/benchmarks/concat_slices	5.713s
```
