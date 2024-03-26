package main

const Infinity = int(^uint(0) >> 1)

func removeNode(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func dijkstra(graph *Graph, start, end int) (int, []int) {
	dist := make(map[int]int)
	prev := make(map[int]int)
	unvisited := make([]int, 0, len(graph.vertices))

	for v := range graph.vertices {
		dist[v] = Infinity
		prev[v] = -1
		unvisited = append(unvisited, v)
	}

	dist[start] = 0

	for len(unvisited) > 0 {
		minVertex := -1

		// Find the unvisited vertex with the smallest distance
		for _, v := range unvisited {
			if minVertex == -1 || dist[v] < dist[minVertex] {
				minVertex = v
			}
		}

		// If we didn't find a vertex with finite distance, we're done
		if minVertex == -1 {
			break
		}

		unvisited = removeNode(unvisited, minVertex)

		// Update the distances to the neighbors of the current vertex
		for _, edge := range graph.vertices[minVertex].edges {
			alt := dist[minVertex] + edge.weight

			if alt < dist[edge.vertex.val] {
				dist[edge.vertex.val] = alt
				prev[edge.vertex.val] = minVertex
			}
		}
	}

	path := []int{}
	curr := end

	for curr != -1 && curr != start {
		path = append([]int{curr}, path...)
		curr = prev[curr]
	}

	if curr == start {
		path = append([]int{start}, path...)
	}

	return dist[end], path
}
