package main

import "math/rand/v2"

type Graph struct {
	vertices map[int]*Vertex
}

type Vertex struct {
	val   int
	edges map[int]*Edge
}

type Edge struct {
	weight int
	vertex *Vertex
}

func (g *Graph) addVertex(val int) {
	if g.vertices == nil {
		g.vertices = make(map[int]*Vertex)
	}

	g.vertices[val] = &Vertex{val: val, edges: make(map[int]*Edge)}
}

func (v *Vertex) addEdge(destination *Vertex, weight int) {
	v.edges[destination.val] = &Edge{weight: weight, vertex: destination}
}

func randomGraph(vertices, edges, maxWeight int) *Graph {
	graph := &Graph{}

	for i := range vertices {
		graph.addVertex(i)
	}

	for _ = range edges {
		source := rand.IntN(vertices)
		destination := rand.IntN(vertices)
		weight := rand.IntN(maxWeight) + 1

		graph.vertices[source].addEdge(graph.vertices[destination], weight)
		graph.vertices[destination].addEdge(graph.vertices[source], weight)
	}

	return graph
}
