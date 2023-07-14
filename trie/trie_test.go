package trie

import (
	"testing"
)

var t = &trie[int]{}

func init() {
	t.Insert("/hello/:username/:addr/", 0)
	t.Insert("/hello/:username/:adda", -111)
	t.Insert("/hello/:username/:adda/b", 1)
	t.Insert("/hello/:username/:adda/c", 2)
	t.Insert("/hello/:username/:adda/d", 3)
	t.Insert("/hello/:username/:adda/e", 4)
	t.Insert("/hello/:username/:adda/f", 5)
	t.Insert("/a/:1/2/:2/:2", 6)
}

var p = []byte("/hello/1/2")

func BenchmarkMyTireRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Insert("/a/b", 6)
		t.Delete("/a/b")
	}
}

func BenchmarkMyTire(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.GetValue(p)
		// t.Insert("/hello/:username/:addr/", "xixi2")
		// t.ParseParams(p)
	}
}

var n1 = new(node)

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
