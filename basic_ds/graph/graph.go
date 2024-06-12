package graph

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{key: key}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) AddEdge(from, to *Vertex) {
	from.adjacent = append(from.adjacent, to)
	to.adjacent = append(to.adjacent, from)
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %d", v.key)
		for _, e := range v.adjacent {
			fmt.Printf("  %d", e.key)
		}
		fmt.Println()
	}
}
