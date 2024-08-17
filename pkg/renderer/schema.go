package renderer

import "github.com/gusfcarvalho/gmh/pkg/models"

type Renderer interface {
	Render([]*models.Node) ([]byte, error)
}
