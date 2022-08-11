package btrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
)

func TestTree1AreEqual(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(1)
	found := Same(t1, t2)
	assert.True(t, found)
}

func TestTree1Tree2AreNotEqual(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(2)
	found := Same(t1, t2)
	assert.False(t, found)
}

func TestTree2Tree2AreEqual(t *testing.T) {
	t1 := tree.New(2)
	t2 := tree.New(2)
	found := Same(t1, t2)
	assert.True(t, found)
}

func TestTree1Tree3AreNotEqual(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(3)
	found := Same(t1, t2)
	assert.False(t, found)
}
