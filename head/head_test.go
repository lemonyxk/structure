/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-06-30 02:13
**/

package head

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	var h = NewMax(9, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	var s = h.Size()

	var str = ""

	for i := 0; i < s; i++ {
		var v, _ = h.Pop()
		str += fmt.Sprintf("%d", v)
	}

	assert.Equal(t, "15109987654321", str)

	var h1 = NewMax(1)

	var v, b = h1.Pop()

	assert.Equal(t, 1, v)
	assert.Equal(t, b, true)

	var h2 = NewMax[int]()
	var v2, b2 = h2.Pop()
	assert.Equal(t, 0, v2)
	assert.Equal(t, b2, false)
}

func TestMin(t *testing.T) {
	var h = NewMin(9, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	var s = h.Size()

	var str = ""

	for i := 0; i < s; i++ {
		var v, _ = h.Pop()
		str += fmt.Sprintf("%d", v)
	}

	assert.Equal(t, "12345678991015", str)

	var h1 = NewMin(1)

	var v, b = h1.Pop()

	assert.Equal(t, 1, v)
	assert.Equal(t, b, true)

	var h2 = NewMin[int]()
	var v2, b2 = h2.Pop()
	assert.Equal(t, 0, v2)
	assert.Equal(t, b2, false)
}
