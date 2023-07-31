/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2023-08-01 00:05
**/

package trie

import (
	"fmt"
	"sync"
)

var lock = new(sync.RWMutex)

type Node[T any] struct {
	Children    map[string]*Node[T]
	Path        string
	Data        T
	Placeholder string
	HasValue    bool
}

func New[T any]() *Node[T] {
	return &Node[T]{
		Children: make(map[string]*Node[T]),
	}
}

func (n *Node[T]) Insert(path string, data T) {

	lock.Lock()
	defer lock.Unlock()

	var pathArray = Split(path)

	if len(pathArray) > 128 {
		panic("path is too long")
	}

	if len(pathArray) == 0 {
		panic("path is empty")
	}

	var node = n

	for i := 0; i < len(pathArray); i++ {
		var out = false

		if pathArray[i] == "" {
			continue
		}

		if node.Children == nil {
			node.Children = make(map[string]*Node[T])
		}

		switch pathArray[i][0] {
		case ':':
			if _, ok := node.Children[":"]; !ok {
				node.Children[":"] = &Node[T]{
					Placeholder: pathArray[i][1:],
				}
			}

			node = node.Children[":"]

		case '*':
			if _, ok := node.Children["*"]; !ok {
				node.Children["*"] = &Node[T]{}
			}

			node = node.Children["*"]

			out = true // * must be the last one
		default:
			if _, ok := node.Children[pathArray[i]]; !ok {
				node.Children[pathArray[i]] = &Node[T]{}
			}

			node = node.Children[pathArray[i]]
		}

		if out {
			break
		}
	}

	if node.HasValue {
		panic(fmt.Sprintf("path %s is conflict with %s", path, node.Path))
	}

	node.HasValue = true
	node.Data = data
	node.Path = path
}

func (n *Node[T]) ParseParams(path string) map[string]string {

	var result = make(map[string]string)

	var pathArray = Split(path)

	var nodePathArray = Split(n.Path)

	if len(pathArray) != len(nodePathArray) {
		return result
	}

	for i := 0; i < len(pathArray); i++ {
		if pathArray[i] == "" {
			continue
		}

		if nodePathArray[i] == "" {
			continue
		}

		if nodePathArray[i][0] == ':' {
			result[nodePathArray[i][1:]] = pathArray[i]
		}
	}

	return result
}

func (n *Node[T]) GetValue(path string) *Node[T] {

	lock.RLock()
	defer lock.RUnlock()

	var pathArray = Split(path)

	var node = n

	var fail *Node[T]

	for i := 0; i < len(pathArray); i++ {
		if pathArray[i] == "" {
			continue
		}

		if node.Children == nil {
			break
		}

		var _, ok = node.Children[pathArray[i]]
		if ok {
			node = node.Children[pathArray[i]]
			continue
		} else {
			// if not found, try to find *
			var _, ok1 = node.Children["*"]
			if ok1 {
				// if found, set fail to *
				if fail == nil { // only set once
					fail = node.Children["*"]
				}
			}
		}

		// if not found, try to find :
		var _, ok2 = node.Children[":"]
		if ok2 {
			node = node.Children[":"]
			continue
		}
	}

	if !node.HasValue { // if not found, return fail
		return fail
	}

	return node
}

func (n *Node[T]) GetAllValue() []*Node[T] {

	lock.RLock()
	defer lock.RUnlock()

	var result []*Node[T]

	var fn func(node *Node[T])

	fn = func(node *Node[T]) {
		if node.Children == nil {
			result = append(result, node)
			return
		}

		for _, value := range node.Children {
			fn(value)
		}
	}

	fn(n)

	return result
}

var pathArr = make([]string, 0, 128)

func Split(path string) []string {
	pathArr = pathArr[:0]

	var index = 0

	for i, v := range path {
		if v == '/' && i == 0 {
			pathArr = append(pathArr, "/")
			index = i
			continue
		}

		if v == '/' && i != 0 {
			if i-index > 1 {
				pathArr = append(pathArr, path[index+1:i])
			}
			pathArr = append(pathArr, "/")
			index = i
		}

		if i == len(path)-1 {
			if i-index > 0 {
				pathArr = append(pathArr, path[index+1:])
			}
		}
	}

	return pathArr
}
