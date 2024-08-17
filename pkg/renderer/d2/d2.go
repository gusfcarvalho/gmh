package d2

import (
	"context"

	"github.com/gusfcarvalho/gmh/pkg/models"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2lib"
)

type D2Render struct {
	graph *d2graph.Graph
}

func (d *D2Render) New() {
	ctx := context.Background()
	// Start with a new, empty graph
	_, graph, err := d2lib.Compile(ctx, "", nil, nil)
	if err != nil {
		panic(err)
	}
	d.graph = graph
}
func (d *D2Render) Render([]*models.Node) ([]byte, error) {
	return nil, nil
}
