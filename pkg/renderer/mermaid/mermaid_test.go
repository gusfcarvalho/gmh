package mermaid_test

import (
	"testing"

	"github.com/gusfcarvalho/gmh/pkg/models"
	"github.com/gusfcarvalho/gmh/pkg/renderer/mermaid"
)

func TestRenderMermaidTD(t *testing.T) {
	nodes := []*models.Node{
		{
			ID:    "A",
			Label: "Node A",
			Children: []*models.Node{
				{
					ID:    "B",
					Label: "Node B",
				},
				{
					ID:    "C",
					Label: "Node C",
				},
			},
		},
		{
			ID:    "D",
			Label: "Node D",
			Children: []*models.Node{
				{
					ID:    "E",
					Label: "Node E",
				},
			},
		},
	}

	expected := `graph TD
  A[Node A]
  A --> B
  A --> C
  D[Node D]
  D --> E
`

	result := mermaid.RenderMermaidTD(nodes)
	if result != expected {
		t.Errorf("RenderMermaidTD() = %s, expected %s", result, expected)
	}
}
