// Package processor implements the core transformation processing logic.
package processor

import (
	"errors"
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/cekrem/go-transform/pkg/domain"
)

var (
	// ErrPluginInterface indicates that a plugin doesn't implement the required interface.
	ErrPluginInterface = errors.New("plugin does not implement TransformerPlugin interface")
	// ErrTransformerNotFound indicates that the requested transformer wasn't found.
	ErrTransformerNotFound = errors.New("transformer not found")
)

// Processor manages the loading and execution of transformation plugins.
type Processor struct {
	plugins map[string]domain.Plugin
}

// NewProcessor creates and initializes a new Processor instance.
func NewProcessor() *Processor {
	return &Processor{
		plugins: make(map[string]domain.Plugin),
	}
}

// LoadPlugin loads a plugin from the given path and registers it with the processor.
func (p *Processor) LoadPlugin(path string) error {
	plug, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open plugin: %w", err)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return fmt.Errorf("plugin does not export 'Plugin': %w", err)
	}

	transformerPlugin, ok := symPlugin.(domain.Plugin)
	if !ok {
		return ErrPluginInterface
	}

	// Use the filename without extension as the plugin name
	name := filepath.Base(path[:len(path)-3]) // remove .so
	p.plugins[name] = transformerPlugin
	return nil
}

// Process executes the named transformer on the input data.
func (p *Processor) Process(transformerName string, input []byte) ([]byte, error) {
	plug, exists := p.plugins[transformerName]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrTransformerNotFound, transformerName)
	}

	transformer := plug.NewTransformer()
	return transformer.Transform(input)
}
