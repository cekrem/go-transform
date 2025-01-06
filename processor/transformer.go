package processor

// Moving from internal/core/ports to processor package
type Transformer interface {
	Transform(input []byte) ([]byte, error)
}

type TransformerPlugin interface {
	NewTransformer() Transformer
}
