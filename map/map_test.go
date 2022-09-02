/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-06-30 02:30
**/

package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	var m = New[int, int]()

	m.Set(0, 0)

	m.Set(1, 1)

	assert.Equal(t, 2, m.Len())

	assert.Equal(t, 0, m.Get(0))
	assert.Equal(t, 1, m.Get(1))

	m.Delete(0)

	assert.Equal(t, 1, m.Len())

	m.Delete(1)

	assert.Equal(t, 0, m.Len())
}

func TestSync(t *testing.T) {
	var m = NewSync[string, string]()

	m.Set("hello", "world")
	m.Set("hello1", "world1")

	assert.Equal(t, "world", m.Get("hello"))
	assert.Equal(t, "world1", m.Get("hello1"))
	assert.Equal(t, "", m.Get("none"))
}
