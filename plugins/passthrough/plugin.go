package main

import "github.com/cekrem/go-transform/pkg/interfaces"

type passthrough struct{}

func (p *passthrough) Transform(input []byte) ([]byte, error) {
	return input, nil
}

type plugin struct{}

func (p *plugin) NewTransformer() interfaces.Transformer {
	return &passthrough{}
}

// Plugin exported symbol
var Plugin plugin
