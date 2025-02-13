// Package main implements a passthrough transformer plugin that returns input without modification.
package main

import "github.com/cekrem/go-transform/pkg/domain"

// passthroughPlugin implements transformer.Plugin without requiring any state.
type passthroughPlugin struct{}

// NewTransformer returns a new passthrough transformer instance.
func (*passthroughPlugin) NewTransformer() domain.Transformer {
	return &passthroughTransformer{}
}

// passthroughTransformer implements transformer.Transformer without requiring any state.
type passthroughTransformer struct{}

// Transform implements transformer.Transformer by returning the input bytes unmodified.
func (pt *passthroughTransformer) Transform(input []byte) ([]byte, error) {
	return input, nil
}

// Plugin exports the passthrough transformer plugin for dynamic loading.
var Plugin passthroughPlugin
var _ = Plugin

// Compile-time verification that passthroughPlugin implements domain.Plugin
var _ domain.Plugin = (*passthroughPlugin)(nil)
