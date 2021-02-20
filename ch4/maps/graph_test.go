package maps

import (
	"fmt"
	"testing"
)

func TestGraphEdge(t *testing.T) {
	addEdge("a", "b")

	if !hasEdge("a", "b") {
		t.Fail()
	}

	fmt.Println(graph)
}
