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

	hash "github.com/lemonyxk/structure/map"
)

func main() {
	var h = hash.New[string, string]()

	h.Set("1", "a")

	log.Println(h.Get("1"))

	var s = hash.NewSync[int, int]()

	s.Set(1, 1)

	// s.Range(func[K, V any](k K, v V) bool {
	// 	log.Println(k, v)
	// 	return true
	// })

	log.Println(s.Get(1))

	s.Range(func(k int, v int) bool {
		log.Println(k, v)
		return true
	})
}
