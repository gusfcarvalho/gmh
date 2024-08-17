package d2

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gusfcarvalho/gmh/pkg/models"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type D2Render struct {
	graph *d2graph.Graph
	ruler *textmeasure.Ruler
}

func (d *D2Render) New() {
	ctx := context.Background()
	// Start with a new, empty graph
	_, graph, err := d2lib.Compile(ctx, "", nil, nil)
	// Initialize a ruler to measure glyphs of text
	d.ruler, _ = textmeasure.NewRuler()
	if err != nil {
		panic(err)
	}
	d.graph = graph
}
func (d *D2Render) Render(nodes []*models.Node) ([]byte, error) {
	for _, node := range nodes {
		// Add a node to the graph
		graph, _, err := d2oracle.Create(d.graph, nil, node.ID)
		if err != nil {
			return nil, err
		}
		d.graph = graph
		for _, child := range node.Children {
			graph, _, err := d2oracle.Create(d.graph, nil, fmt.Sprintf("%s.%s", node.ID, child.ID))
			if err != nil {
				return nil, err
			}
			d.graph = graph
		}
	}
	script := d2format.Format(d.graph.AST)
	// Compile the script with given theme and layout
	diagram, _, err := d2lib.Compile(context.Background(), script, &d2lib.CompileOptions{
		Ruler: d.ruler,
		LayoutResolver: func(engine string) (d2graph.LayoutGraph, error) {
			return d2dagrelayout.DefaultLayout, nil
		},
	}, nil)
	if err != nil {
		return nil, err
	}
	// Render to SVG
	padding := int64(d2svg.DEFAULT_PADDING)
	out, _ := d2svg.Render(diagram, &d2svg.RenderOpts{
		Pad: &padding,
	})

	// Write to disk
	_ = os.WriteFile(filepath.Join("svgs", "out.svg"), out, 0600)
	return nil, nil
}
