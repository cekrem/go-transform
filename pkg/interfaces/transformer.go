// Package interfaces provides core abstractions for transformation operations.
package interfaces

// Transformer defines the interface for data transformation operations.
type Transformer interface {
	// Transform processes the input bytes and returns transformed bytes or an error.
	Transform(input []byte) ([]byte, error)
}

// TransformerPlugin defines the interface for plugin implementations.
type TransformerPlugin interface {
	// NewTransformer creates and returns a new Transformer instance.
	NewTransformer() Transformer
}
