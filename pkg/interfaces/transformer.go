package interfaces

// Core abstractions that other packages depend on.
type Transformer interface {
	Transform(input []byte) ([]byte, error)
}

type TransformerPlugin interface {
	NewTransformer() Transformer
}
