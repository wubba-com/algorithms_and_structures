package graph

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	var err error
	// граф              (7)             (1)
	//         B (1) --------- F (5) ---------- G (6)
	//       /               /    |  \
	//     /  (2)          / (2)  |   \ (4)
	//   /               /        |    \
	// A (0)           D (3)      | (1)  S (7)
	//  \            /            |    /
	//   \ (1)     / (5)          |  / (6)
	//    \      /     (2)        |/
	//     C (2) ------------- E (4)
	// матрица смежности графа:

	//   0 1 2 3 4 5 6
	// 0 x 1 1 x x x x
	// 1 x x x x 1 x x
	// 2 x x x 1 x 1 x
	// 3 x x x x 1 x x
	// 4 x x x x x x 1
	// 5 x x x x x x 1
	// 6 x x x x x x x
	g := New()
	idA := g.AddNode()
	idB := g.AddNode()
	idC := g.AddNode()
	idD := g.AddNode()
	idE := g.AddNode()
	idF := g.AddNode()
	idG := g.AddNode()
	idS := g.AddNode()

	err = g.AddEdge(idA, idB, 2)
	assert.NoError(t, err)
	err = g.AddEdge(idB, idA, 2)
	assert.NoError(t, err)
	err = g.AddEdge(idA, idC, 1)
	assert.NoError(t, err)
	err = g.AddEdge(idB, idF, 7)
	assert.NoError(t, err)
	err = g.AddEdge(idC, idD, 5)
	assert.NoError(t, err)
	err = g.AddEdge(idC, idE, 2)
	assert.NoError(t, err)
	err = g.AddEdge(idD, idF, 2)
	assert.NoError(t, err)
	err = g.AddEdge(idE, idF, 1)
	assert.NoError(t, err)
	err = g.AddEdge(idF, idG, 1)
	assert.NoError(t, err)
	err = g.AddEdge(idF, idS, 4)
	assert.NoError(t, err)
	err = g.AddEdge(idE, idS, 6)
	assert.NoError(t, err)

	exist := g.BreadthSearch(idA, idS)
	assert.Equal(t, true, exist)

	costs, path := g.Short(idA, idS)
	assert.Equal(t, 8, costs)
	reflect.DeepEqual([]int{0, 2, 4, 5, 7}, path)

}
