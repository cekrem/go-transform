package domain

// Plugin defines the interface for plugin implementations.
type Plugin interface {
	// NewTransformer creates and returns a new Transformer instance.
	NewTransformer() Transformer
}
