package processor

import (
	"errors"
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/cekrem/go-transform/pkg/interfaces"
)

var (
	ErrPluginInterface     = errors.New("plugin does not implement TransformerPlugin interface")
	ErrTransformerNotFound = errors.New("transformer not found")
)

type Processor struct {
	plugins map[string]interfaces.TransformerPlugin
}

func NewProcessor() *Processor {
	return &Processor{
		plugins: make(map[string]interfaces.TransformerPlugin),
	}
}

func (p *Processor) LoadPlugin(path string) error {
	plug, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open plugin: %w", err)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return fmt.Errorf("plugin does not export 'Plugin': %w", err)
	}

	transformerPlugin, ok := symPlugin.(interfaces.TransformerPlugin)
	if !ok {
		return fmt.Errorf("plugin does not implement TransformerPlugin interface")
	}

	// Use the filename without extension as the plugin name
	name := filepath.Base(path[:len(path)-3]) // remove .so
	p.plugins[name] = transformerPlugin
	return nil
}

func (p *Processor) Process(transformerName string, input []byte) ([]byte, error) {
	plugin, exists := p.plugins[transformerName]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrTransformerNotFound, transformerName)
	}

	transformer := plugin.NewTransformer()
	return transformer.Transform(input)
}
