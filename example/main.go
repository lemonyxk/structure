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

	t.Insert("/hello/:username/:addr/", "xixi2")
	t.Insert("/hello/:username/:adda", "xixi1")
	t.Insert("/hello/:username/:adda/b", 1)
	t.Insert("/hello/:username/:adda/c", 2)
	t.Insert("/hello/:username/:adda/d", 3)
	t.Insert("/hello/:username/:adda/e", 4)
	t.Insert("/hello/:username/:adda/f", 5)
	t.Insert("/a/:1/2/:2/:2", 6)

	var p = []byte("/a/1/2/1/ddsadasdsadsadsdsadsaddsadasdsadsadsdsadsa")

	if t.GetValue(p) != nil {
		log.Println(t.GetValue(p).ParseParams(p))
		log.Println(string(t.GetValue(p).Path), t.GetValue(p).Keys)
		log.Println([]byte("*"))
	}

}
