# allocations
Understanding the differences between the stack and the heap.

## How to Use It

### View compiler output
```
go build -gcflags '-m -l'
```

### Run benchmarks
```
go test -bench . -benchmem
```

- Reading banechmark result:
    ```shell

    goos: darwin                                        # My system OS
    goarch: amd64                                       # My system arch
    pkg: github.com/Jimeux/go-samples/allocations       # package or go.mod module
    cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz       # My CPU

  | benchFuncName-NoOfCPUs             | Noofinteration    | AvgTimeEachIteration | BytesPerOp    | AllocsPerOp |
    BenchmarkStackIt_Invoke_Main1-12     	1000000000	       1.050 ns/op	          0 B/op          0 allocs/op
    BenchmarkStackIt_Invoke_Main2-12     	95918215	         12.07 ns/op	          8 B/op	        1 allocs/op
    BenchmarkStackIt_Invoke_Main3-12    	953234210	         1.244 ns/op	          0 B/op          0 allocs/op
    Benchmark_Invoke_CopyIt-12           	251543496	         4.686 ns/op	          0 B/op          0 allocs/op
    Benchmark_Invoke_PointerIt-12        	29533814	         37.86 ns/op	          80 B/op	        1 allocs/op

    ```


### Generate & view `CreateCopy` trace
```
go test -run TestCopyIt -trace=copy_trace.out
go tool trace copy_trace.out
```

### Generate & view `CreatePointer` trace
```
go test -run TestPointerIt -trace=copy_trace.out
go tool trace pointer_trace.out
```
