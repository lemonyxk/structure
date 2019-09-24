    package main
    
    import (
    	"log"
    	
    	"tire"
    )
    
    func init() {
    	log.SetFlags(log.Lshortfile)
    }
    
    func main() {
    	
    	var t = &tire.Tire{}
    	
    	t.Insert("/:username/lemo/:addr/", "xixi2")
    	t.Insert("/hell/:username/:adda/c/:xixi/:haha", 2)
    	t.Insert("/hello/:username/:adda/aa", "xixi1")
    	t.Insert("/hello/:username/:adda/b", 1)
    	t.Insert("/hello/:username/:adda/d", 3)
    	t.Insert("/hello/:username/:adda/e", 4)
    	t.Insert("/hello/:username/:adda/:f", 5)
    	t.Insert("/:1/:2", 6)
    	t.Insert("/a/:1/2/:2/:2", 6)
    	
    	var p = []byte("/hello/lemo/addr/f")
    	
    	log.Println(string(t.GetValue(p).Path))
    	
    }
    
    
    
    
    
    short path: /a/1/2/1/d
    
    ⇒  go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out -count=5
    goos: darwin
    goarch: amd64
    pkg: tire
    BenchmarkMyTire-4       41385664                24.6 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       43024808                24.5 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       47878970                24.6 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       43786089                24.5 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       48587972                24.5 ns/op             0 B/op          0 allocs/op
    BenchmarkTireTest-4      9321482               121 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      8969742               123 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      9547354               122 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      9446014               122 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      9755164               121 ns/op              96 B/op          1 allocs/op
    PASS
    ok      tire    13.193s
    
    
    
    
    long path: /a/1/2/1/ddsadasdsadsadsdsadsaddsadasdsadsadsdsadsa
    
    ⇒  go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out -count=5
    goos: darwin
    goarch: amd64
    pkg: tire
    BenchmarkMyTire-4       12776796                87.1 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       13063708                88.5 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       13336400                87.4 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       13413303                86.0 ns/op             0 B/op          0 allocs/op
    BenchmarkMyTire-4       13509172                85.9 ns/op             0 B/op          0 allocs/op
    BenchmarkTireTest-4      8193919               141 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      8187244               145 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      7447316               144 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      8175390               174 ns/op              96 B/op          1 allocs/op
    BenchmarkTireTest-4      7655865               229 ns/op              96 B/op          1 allocs/op
    PASS
    ok      tire    13.718s

