/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-12-26 19:12
**/

package tire

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

const FH uint8 = 58
const XG uint8 = 47
const SC uint8 = 255
const WH uint8 = 63

var mux sync.Mutex

func New[T any]() *Tire[T] {
	return &Tire[T]{}
}

type Tire[T any] struct {
	children      *[SC]*Tire[T]
	parent        *Tire[T]
	char          byte
	childrenCount uint8
	Keys          []string
	Path          []byte
	Data          T
}

func (t *Tire[T]) ParseParams(pathBytes []byte) []string {

	if len(t.Keys) == 0 {
		return nil
	}

	var pathArray = strings.Split(bytesToString(pathBytes), string(XG))

	var res []string

	var i = 1

	var bLen = len(t.Path) - 1

	for index, value := range t.Path {

		if value == XG && index != bLen && t.Path[index+1] == FH {
			res = append(res, pathArray[i])
			i++
			continue
		}

		if value == XG {
			i++
		}
	}

	return res
}

func (t *Tire[T]) Insert(path string, data T) {

	mux.Lock()
	defer mux.Unlock()

	if path == "" {
		panic("path is empty")
	}

	var pathBytes = stringToBytes(path)

	if pathBytes[0] != XG {
		panic(fmt.Sprintf("path must start with [%s]", string(XG)))
	}

	// include ?
	if strings.Index(path, string(WH)) != -1 {
		panic(fmt.Sprintf("%s is include [%s]", path, string(WH)))
	}

	// repeat router
	if h := getFormatValue(t, formatPath(pathBytes)); h != nil {
		panic(fmt.Sprintf("%s is conflict with %s", path, h.Path))
	}

	var t1 = t

	var k []byte

	var ka []string

	var s = true

	for index := range pathBytes {

		var c = pathBytes[index]

		// if c > SC {
		//	panic(fmt.Sprintf("%s include special characters %s", path, string(c)))
		// }

		if c == FH && (index != 0 && pathBytes[index-1] == XG) {
			s = false
			if index == len(pathBytes)-1 || (index != 0 && pathBytes[index+1] == XG) {
				panic(fmt.Sprintf("%s is invalid, after [:] do not have any var", path))
			}
			continue
		}

		if s == false {
			k = append(k, c)
			if index == len(pathBytes)-1 || (index != 0 && pathBytes[index+1] == XG) {
				c = FH
			} else {
				continue
			}
		}

		if s == false {
			c = FH
		}

		if k != nil {
			ka = append(ka, bytesToString(k))
		}

		var p *Tire[T]

		if t.children == nil {
			t.children = &[SC]*Tire[T]{}
		}

		if t.children[c] != nil {
			p = t.children[c]
		} else {
			p = new(Tire[T])
			p.parent = t
			p.children = &[SC]*Tire[T]{}
			p.char = c
		}

		if index == len(pathBytes)-1 {
			p.Keys = ka
			p.Path = pathBytes
			p.Data = data
		}

		t.children[c] = p
		t.childrenCount++

		t = p
		k = nil
		s = true

	}

	t = t1

}

func formatPath(pathBytes []byte) []byte {

	if pathBytes == nil || len(pathBytes) == 0 {
		return nil
	}

	if pathBytes[0] != XG {
		return nil
	}

	if len(pathBytes) == 1 {
		return []byte{XG}
	}

	var res []byte

	var s = true

	for index := range pathBytes {
		var c = pathBytes[index]

		if c == FH && pathBytes[index-1] == XG {
			res = append(res, c)
			s = false
			continue
		}

		if c == XG {
			s = true
		}

		if s == true {
			res = append(res, pathBytes[index])
		}

	}

	return res
}

func getFormatValue[T any](t *Tire[T], pathBytes []byte) *Tire[T] {

	var n = t.children

	if t.childrenCount == 0 {
		return nil
	}

	for index := range pathBytes {

		var c = pathBytes[index]

		if n[c] == nil {
			return nil
		}

		if n[c].char != 0 {

			if index == len(pathBytes)-1 && n[c].Path != nil {
				return n[c]
			}

			n = n[c].children
		}

	}

	return nil

}

func (t *Tire[T]) Delete(path string) {

	mux.Lock()
	defer mux.Unlock()

	var pathBytes = stringToBytes(path)

	var node = t.GetValue(pathBytes)

	if node == nil {
		return
	}

	for {

		node.parent.children[node.char] = nil

		node = node.parent

		if node.childrenCount != 0 {
			break
		}
	}
}

func (t *Tire[T]) GetValue(pathBytes []byte) *Tire[T] {

	var n = t.children

	if t.childrenCount == 0 {
		return nil
	}

	var bLen = len(pathBytes) - 1

	var f = true

	for index := range pathBytes {

		// c == : ?

		var c = pathBytes[index]

		if c == FH {

			// is the latest char ?
			if index == bLen {

				if n[FH] != nil && n[FH].Path != nil {
					return n[FH]
				}

				return nil
			}

			continue
		}

		if n[c] == nil || f == false {

			f = false

			// is /
			if c == XG {

				// is the latest char ?
				if index == bLen {

					if n[FH] == nil {
						return nil
					}

					if n[FH].children[XG] == nil {
						return nil
					}

					if n[FH].children[XG].Path == nil {
						return nil
					}

					// remove / index
					return n[FH].children[XG]
				}

				// not the latest char
				// return nil
				if n[FH] == nil {
					return nil
				}

				// not children return nil
				if n[FH].children[XG] == nil {
					return nil
				}

				// reset n
				n = n[FH].children[XG].children

				f = true

				continue

			}

			// is the latest char ?
			if index == bLen {
				if n[FH] != nil && n[FH].Path != nil {
					return n[FH]
				}

				return nil
			}

			if n[FH] == nil {
				return nil
			}

			continue
		}

		if n[c].char != 0 {
			if index == bLen && n[c].Path != nil {
				return n[c]
			}

			n = n[c].children

			f = true
		}
	}

	return nil
}

func fn[T any](node *Tire[T], res *[]*Tire[T]) {
	if node == nil {
		return
	}
	for i := 0; i < len(node.children); i++ {
		if node.children[i] != nil {
			if len(node.children[i].Path) != 0 {
				*res = append(*res, node.children[i])
			}
			fn(node.children[i], res)
		}
	}
}

func stringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (t *Tire[T]) GetAllValue() []*Tire[T] {
	var res []*Tire[T]
	fn(t, &res)
	return res
}
