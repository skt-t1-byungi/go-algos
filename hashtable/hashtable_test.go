package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	ht := New()

	ht.Set("a", 10)
	v, _ := ht.Get("a")
	assert.Equal(t, 10, v)

	ht.Set("b", 2)
	v, _ = ht.Get("b")
	assert.Equal(t, 2, v)

	assert.Equal(t, []string{"a", "b"}, ht.Keys())
	assert.Equal(t, []interface{}{10, 2}, ht.Values())
}
