// Package main implements a passthrough transformer plugin.
package main

import "github.com/cekrem/go-transform/pkg/interfaces"

// passthroughTransformer implements a transformer that returns input unchanged.
type passthroughTransformer struct{}

// Transform returns the input bytes without modification.
func (pt *passthroughTransformer) Transform(input []byte) ([]byte, error) {
	return input, nil
}

// passthroughPlugin implements the TransformerPlugin interface.
type passthroughPlugin struct{}

// NewTransformer creates a new passthrough transformer instance.
func (*passthroughPlugin) NewTransformer() interfaces.Transformer {
	return &passthroughTransformer{}
}

// Plugin is the exported symbol required by the plugin system.
var Plugin passthroughPlugin
