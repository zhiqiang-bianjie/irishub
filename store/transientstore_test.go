package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var k, v = []byte("hello"), []byte("world")

func TestTransientStore(t *testing.T) {
	tstore := newTransientStore()

	require.Nil(t, tstore.Get(k))

	tstore.Set(k, v)

	require.Equal(t, v, tstore.Get(k))

	tstore.Commit(nil)

	require.Nil(t, tstore.Get(k))
}
