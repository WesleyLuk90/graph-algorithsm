package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToGraph(t *testing.T) {
	g := toGraph("a,b;c,d")
	assert.Equal(t, len(g.Edges), 2)
	assert.Equal(t, g.Edges[0].A, "a")
	assert.Equal(t, g.Edges[0].B, "b")
	assert.Equal(t, g.Edges[1].A, "c")
	assert.Equal(t, g.Edges[1].B, "d")
}

func TestVertexCover(t *testing.T) {
	g := toGraph("a,b;b,c;c,d;a,d")
	vcs := GenerateVertexCover(g)
	assert.Equal(t, len(vcs), 16)
}

func TestCovers(t *testing.T) {
	g := toGraph("a,b;b,c;c,d;a,d")
	vc := VertexCover{[]string{"a", "c"}}
	assert.True(t, vc.Covers(g))
}