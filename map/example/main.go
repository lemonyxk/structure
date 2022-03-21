/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-03-21 23:44
**/

package main

import (
	"log"

	hash "github.com/lemonyxk/structure/v3/map"
)

func main() {
	var h = hash.New[string, string]()

	h.Set("1", "a")

	log.Println(h.Get("1"))
}
