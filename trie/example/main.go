package main

import (
	// "github.com/lemoyxk/structure/trie"

	"github.com/lemonyxk/structure/trie"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	var t = trie.New[int]()

	t.Insert("/resource/:type/:hash/:name", -111111111)
	t.Insert("/resource/:type/delete", -2)
	t.Insert("/hello/:username/:adda", -111)
	t.Insert("/hello/:username/:adda/b", 1)
	t.Insert("/hello/:username/:adda/c", 2)
	t.Insert("/hello/:username/:adda/d", 3)
	t.Insert("/hello/:username/:adda/e", 4)
	t.Insert("/hello/:username/:adda/f", 5)
	t.Insert("/a/b/*", 6)
	t.Insert("/a/b/c", 7)
	t.Insert("/a/b/:d", 8)

	var v = t.GetAllValue()
	for _, value := range v {
		log.Println(value.Path, value.Data)
	}

	var res = t.GetValue("/resource/image/d77f9f2aa6b1617e/origin.jpeg")

	log.Println(res.Data)
	log.Println(res.ParseParams("/resource/image/d77f9f2aa6b1617e/origin.jpeg"))

}
