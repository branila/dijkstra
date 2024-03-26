package main

import (
	"fmt"
	"sort"
)

func printGraph(g *Graph) {
	keys := []int{}

	for key := range g.vertices {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		vertex := g.vertices[key]

		fmt.Printf("N%d ->\t", vertex.val)

		for _, edge := range vertex.edges {
			fmt.Printf("%d [%d]\t", edge.vertex.val, edge.weight)
		}

		fmt.Println()
	}

	fmt.Println()
}
