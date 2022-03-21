package main

import (
	// "github.com/lemoyxk/structure/tire"

	"log"

	"github.com/lemoyxk/structure/v3/tire"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	var t = &tire.Tire[int]{}

	t.Insert("/hello/:username/:addr/", -1)
	t.Insert("/hello/:username/:adda", 4)
	t.Insert("/hello/:username/:adda/b", 1)
	t.Insert("/hello/:username/:adda/c", 2)
	t.Insert("/hello/:username/:adda/d", 3)
	t.Insert("/hello/:username/:adda/e", 4)
	t.Insert("/hello/:username/:adda/f", 5)
	t.Insert("/a/:1/2/:2/:2", 6)
	t.Insert("/b/:1/2/:2/:2", 7)
	t.Insert("/c/:1/2/:2/:2", 9)
	t.Insert("/*(21", 9)

	log.Println(t.GetValue([]byte("/hello/:username/:addr/")).Data)

	var a = t.GetAllValue()
	for i := 0; i < len(a); i++ {
		log.Println(a[i].Data, string(a[i].Path))
	}

	//
	// if t.GetValue(p) != nil {
	//	log.Println(t.GetValue(p).ParseParams(p))
	//	log.Println(string(t.GetValue(p).Path), t.GetValue(p).Keys)
	//	log.Println([]byte("*"))
	// }

	// for _, value := range t.GetAllValue() {
	// 	log.Println(string(value.Path), value.Data)
	// }
	//
	// t.Delete("/hello/:username/:addr/")
	// t.Delete("/hello/:username/:adda")
	// t.Delete("/hello/:username/:adda/b")
	// t.Delete("/hello/:username/:adda/c")
	// t.Delete("/hello/:username/:adda/d")
	// t.Delete("/hello/:username/:adda/e")
	// t.Delete("/hello/:username/:adda/f")
	//
	// log.Println("-----------")
	//
	// for _, value := range t.GetAllValue() {
	// 	log.Println(string(value.Path), value.Data)
	// }
}
