package main

import (
	// "github.com/Lemo-yxk/structure/tire"

	"log"

	"github.com/Lemo-yxk/structure/tire"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	var t = &tire.Tire{}

	t.Insert("/hello/:username/:addr/", -1)
	t.Insert("/hello/:username/:adda", 0)
	t.Insert("/hello/:username/:adda/b", 1)
	t.Insert("/hello/:username/:adda/c", 2)
	t.Insert("/hello/:username/:adda/d", 3)
	t.Insert("/hello/:username/:adda/e", 4)
	t.Insert("/hello/:username/:adda/f", 5)
	t.Insert("/a/:1/2/:2/:2", 6)
	t.Insert("/b/:1/2/:2/:2", 7)
	t.Insert("/c/:1/2/:2/:2", 9)
	t.Insert("/*(我", 9)

	log.Println(t.GetValue([]byte("/*(我")).Data)
	//
	// if t.GetValue(p) != nil {
	//	log.Println(t.GetValue(p).ParseParams(p))
	//	log.Println(string(t.GetValue(p).Path), t.GetValue(p).Keys)
	//	log.Println([]byte("*"))
	// }

	for _, value := range t.GetAllValue() {
		log.Println(string(value.Path), value.Data)
	}
}
