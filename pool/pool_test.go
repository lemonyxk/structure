/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-07-03 00:38
**/

package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name string
}

func TestPool(t *testing.T) {

	var pool = New(Config[*Test]{
		Max: 0,
		Min: 0,
		New: func() *Test {
			return &Test{}
		},
	})

	var p = pool.Get()

	p.Name = "test"

	pool.Put(p)

	var p1 = pool.Get()

	assert.Equal(t, p, p1)
}
