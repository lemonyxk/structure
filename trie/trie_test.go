package trie

import (
	"testing"
)

var n1 = new(node)
var p = []byte("/hello/1/2")

func init() {
	// n1.addRoute("/hello/:username/:addr/hello", "xixi2")
	n1.addRoute("/hello/:username/:adda", "xixi1")
	n1.addRoute("/hello/:username/:adda/b", 1)
	n1.addRoute("/hello/:username/:adda/c", 2)
	n1.addRoute("/hello/:username/:adda/d", 3)
	n1.addRoute("/hello/:username/:adda/e", 4)
	n1.addRoute("/hello/:username/:adda/f", 5)
	n1.addRoute("/a/:1/2/:2/:2", 6)
}

func BenchmarkTireTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n1.getValue(string(p))
	}
}

var n2 = New[int]()

func init() {
	n2.Insert("/resource/:type/:hash/:name", -1)
	n2.Insert("/resource/:type/delete", -2)
	n2.Insert("/hello/:username/:adda", -111)
	n2.Insert("/hello/:username/:adda/b", 1)
	n2.Insert("/hello/:username/:adda/c", 2)
	n2.Insert("/hello/:username/:adda/d", 3)
	n2.Insert("/hello/:username/:adda/e", 4)
	n2.Insert("/hello/:username/:adda/f", 5)
	n2.Insert("/a/b/*", 6)
}

func BenchmarkTireTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n2.GetValue("/a/b/image")
	}
}
