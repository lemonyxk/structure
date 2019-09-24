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
