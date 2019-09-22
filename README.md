# quickSort
quickSort implementation in GO and Python

## Use GO Modules
- go mod init
- go mod tidy

## Build Go Binary
- go build

## Test
go test -bench=.

### Here is the result from the benchmarking test comparing QuickSort and GO's default Sort
```
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/another-maverick/quickSort
BenchmarkQuickSort-4             3000000               583 ns/op
BenchmarkDefaultSort-4           2000000               762 ns/op
PASS
ok      github.com/another-maverick/quickSort   4.657s
```
This proves that QuickSort is faster than GO's  default Sort
