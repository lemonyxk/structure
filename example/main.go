package main

import (
	"log"

	"github.com/Lemo-yxk/tire"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	var t = &tire.Tire{}

	// t.Insert("/:username/lemo/:addr/", "xixi2")
	// t.Insert("/hell/:username/:adda/c/:xixi/:haha", 2)
	// t.Insert("/hello/:username/:adda/aa", "xixi1")
	// t.Insert("/hello/:username/:adda/b", 1)
	// t.Insert("/hello/:username/:adda/d", 3)
	// t.Insert("/hello/:username/:adda/e", 4)
	// t.Insert("/hello/:username/:adda/:f", 5)
	// t.Insert("/:1/:21", 6)
	t.Insert("/debug/pprof/:tip", 6)

	var p = []byte("/debug/pprof/heap")

	log.Println(string(t.GetValue(p).Path))

	log.Println([]byte("*"))

}
