package graph

import (
	"errors"
	"github.com/wubba-com/algorithms_and_structures/algorithms/sorting"
	"github.com/wubba-com/algorithms_and_structures/structure/queue"
)

var (
	nonExistentNode = errors.New("non-existent node")
)

func New() *Graph {
	return &Graph{
		nodes: make([]*Node, 0),
	}
}

const (
	maxWeight = 9999999999999
)

type (
	Graph struct {
		nodes []*Node
	}

	Node struct {
		id    int
		edges map[int]int
	}
)

func (g *Graph) AddNode() int {
	var id = len(g.nodes)
	g.nodes = append(g.nodes, &Node{
		id:    id,
		edges: make(map[int]int),
	})

	return id
}

func (g *Graph) get(id int) *Node {
	if id > len(g.nodes)-1 {
		return nil
	}
	return g.nodes[id]
}

func (g *Graph) AddEdge(idNode1, idNode2, weight int) error {
	if idNode1 > len(g.nodes)-1 {
		return nonExistentNode
	}
	g.nodes[idNode1].edges[idNode2] = weight

	return nil
}

func (g *Graph) BreadthSearch(a, b int) bool {
	q := queue.New()
	q.PushBack(a)
	for q.Len() > 0 {
		current, err := q.PopFront()
		if err != nil {
			return false
		}
		idNode := current.Value().(int)
		node := g.get(idNode)
		if node == nil {
			return false
		}
		if _, ok := node.edges[b]; ok {
			return true
		} else {
			for id, _ := range g.nodes[idNode].edges {
				q.PushBack(id)
			}
		}
	}

	return false
}

func (g *Graph) Short(a, b int) (int, []int) {
	var (
		weights   = make(map[int]int)
		processed = make([]int, 0)
		neighbors = make(map[int]int)
	)
	endNode := g.get(b)
	if endNode == nil {
		return -1, nil
	}

	for id := range g.nodes {
		if id != a {
			startNode := g.get(a)
			if startNode == nil {
				return -1, nil
			}

			weights[id] = maxWeight
			if value, ok := startNode.edges[id]; ok {
				weights[id] = value
			}
		}
	}

	// получаем ребро с самым легким весом
	id := g.findNodeWithLowerWeight(weights, processed)
	for id >= 0 {
		weight := weights[id]
		neighbors = g.nodes[id].edges
		for idNeighbor := range neighbors {
			weightNeighbor := neighbors[idNeighbor]
			newWeight := weight + weightNeighbor
			if newWeight < weights[idNeighbor] {
				weights[idNeighbor] = newWeight
			}
		}
		processed = append(processed, id)
		id = g.findNodeWithLowerWeight(weights, processed)
	}

	path := append([]int{}, b)
	end := b
	var i = len(g.nodes)
	for i > 0 {
		neighborsIds := make(map[int]int, 0)
		for node := range g.nodes {
			if w, ok := g.nodes[node].edges[end]; ok {
				neighborsIds[node] = w
			}
		}

		for neighbor, w := range neighborsIds {
			minW := weights[end]
			nodeW := weights[neighbor]
			if minW-nodeW == w {
				path = append(path, neighbor)
				end = neighbor
			}
		}

		i--
	}

	sorting.Quicksort(path)

	return weights[b], path
}

func (g *Graph) findNodeWithLowerWeight(costs map[int]int, processed []int) int {
	var (
		lowerCost   = maxWeight
		idLowerNode = -1
	)

	for idNode := range costs {
		cost := costs[idNode]
		if cost < lowerCost && !sorting.Includes(processed, idNode) {
			lowerCost = cost
			idLowerNode = idNode
		}
	}

	return idLowerNode
}
