package main

import "fmt"

func main() {
	graph := randomGraph(10, 15, 10)

	fmt.Printf("\nGrafo generato:\n\n")
	printGraph(graph)

	var start, end int

	fmt.Print("Inserire il nodo iniziale: ")
	fmt.Scan(&start)

	fmt.Print("Inserire il nodo finale: ")
	fmt.Scan(&end)

	distance, path := dijkstra(graph, start, end)

	if distance == Infinity {
		fmt.Printf("\nDistanza minima: Infinito\n")
	} else {
		fmt.Printf("\nDistanza minima: %d\n", distance)
	}

	fmt.Printf("Percorso minimo: %v\n", path)

	fmt.Println()
}
