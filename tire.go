/**
* @program: tire-tree
*
* @description:
*
* @author: Lemo-yxk
*
* @create: 2019-09-24 15:42
**/

package tire

import (
	"fmt"
	"strings"
)

const FH uint8 = 58
const XG uint8 = 47
const SC uint8 = 255
const WH uint8 = 63
const XH uint8 = 42

type Tire struct {
	children      *[SC]*Tire
	char          byte
	childrenCount uint8
	Keys          []string
	Path          []byte
	Data          interface{}
}

func (t *Tire) ParseParams(pathBytes []byte) []string {

	if len(t.Keys) == 0 {
		return nil
	}

	var pathArray = strings.Split(string(pathBytes), "/")

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

func (t *Tire) Insert(path string, data interface{}) {

	if path == "" {
		return
	}

	var pathBytes = []byte(path)

	if pathBytes[0] != XG {
		return
	}

	// include ?
	if strings.Index(path, string(WH)) != -1 {
		panic(fmt.Sprintf("%s is include [%s]", path, string(WH)))
	}

	// repeat router
	if h := t.GetValue(pathBytes); h != nil {
		panic(fmt.Sprintf("%s is conflict with %s", path, h.Path))
	}

	var t1 = t

	var k []byte

	var ka []string

	var s = true

	for index := range pathBytes {

		var c = pathBytes[index]

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
			ka = append(ka, string(k))
		}

		var p *Tire

		if t.children == nil {
			t.children = &[SC]*Tire{}
		}

		if t.children[c] != nil {
			p = t.children[c]
		} else {
			p = new(Tire)
			p.children = &[SC]*Tire{}
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

func (t *Tire) GetValue(pathBytes []byte) *Tire {

	var n = t.children

	if t.childrenCount == 0 {
		return nil
	}

	var bLen = len(pathBytes) - 1

	for index := range pathBytes {

		// c == : ?
		if pathBytes[index] == FH {

			// is the latest char ?
			if index == bLen {
				if n[FH] != nil && n[FH].Path != nil {
					return n[FH]
				}

				return nil
			}

			continue
		}

		if n[pathBytes[index]] == nil {

			// is /
			if pathBytes[index] == XG {

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

		if n[pathBytes[index]].char != 0 {
			if index == bLen && n[pathBytes[index]].Path != nil {
				return n[pathBytes[index]]
			}

			n = n[pathBytes[index]].children
		}
	}

	return nil
}
