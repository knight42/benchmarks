goos: darwin
goarch: amd64
pkg: github.com/knight42/benchmarks/map_keys
BenchmarkStrAsKey-4      	  300000	      4359 ns/op	    1162 B/op	      95 allocs/op
BenchmarkIntAsKey-4      	 1000000	      1398 ns/op	     568 B/op	      20 allocs/op
BenchmarkGetStrAsKey-4   	50000000	        36.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetIntAsKey-4   	100000000	        18.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/knight42/benchmarks/map_keys	6.550s
