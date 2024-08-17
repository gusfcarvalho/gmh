// Mermaid renderer for gmh
// This code is licensed under the Apache License, Version 2.0
package mermaid

import (
	"fmt"
	"strings"

	"github.com/gusfcarvalho/gmh/pkg/models"
)

// RenderMermaidTD generates Mermaid TD code from a given []*models.Node
func RenderMermaidTD(nodes []*models.Node) string {
	var sb strings.Builder

	sb.WriteString("graph TD\n")

	for _, node := range nodes {
		sb.WriteString(fmt.Sprintf("  %s[%s]\n", node.ID, node.Label))

		for _, child := range node.Children {
			sb.WriteString(fmt.Sprintf("  %s --> %s\n", node.ID, child.ID))
		}
	}

	return sb.String()
}
