package tfstate

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gusfcarvalho/gmh/pkg/models"
)

type TFStateProvider struct {
}

type State struct {
	Values RootValues `json:"values"`
}
type RootValues struct {
	RootModule Module `json:"root_module"`
}

type Module struct {
	Resources    []Resource `json:"resources"`
	Address      string     `json:"address"`
	ChildModules []Module   `json:"child_modules"`
}

type Resource struct {
	Address       string                 `json:"address"`
	Mode          string                 `json:"mode"`
	Type          string                 `json:"type"`
	Index         any                    `json:"index"`
	Name          string                 `json:"name"`
	ProviderName  string                 `json:"provider_name"`
	SchemaVersion int                    `json:"schema_version"`
	Values        map[string]interface{} `json:"values"`
}

func (f *TFStateProvider) Convert(path string) (*models.Node, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	state := State{}
	err = json.Unmarshal(file, &state)
	if err != nil {
		return nil, err
	}
	state.Values.RootModule.Address = "root"
	nodes := readModule(state.Values.RootModule)
	return nodes, nil
}
func sanitizeID(ID string) string {
	out := strings.ReplaceAll(ID, "\"]", "_")
	out = strings.ReplaceAll(out, "[\"", "_")
	out = strings.ReplaceAll(out, ".", "_")
	out = strings.ReplaceAll(out, "__", "_")
	out = strings.TrimSuffix(out, "_")
	return out

}
func readModule(mod Module) *models.Node {
	module_node := &models.Node{
		ID: sanitizeID(mod.Address),
	}
	for _, res := range mod.Resources {
		address := fmt.Sprintf("%s.%s", res.Type, res.Name)
		if res.Index != nil {
			address = fmt.Sprintf("%s.%s.%s", res.Type, res.Name, res.Index)
		}
		module_node.Children = append(module_node.Children, &models.Node{
			ID: sanitizeID(address),
		})
	}
	for _, child := range mod.ChildModules {
		child_module := readModule(child)
		module_node.Children = append(module_node.Children, child_module)
	}
	return module_node
}
