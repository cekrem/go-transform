// Package domain provides core abstractions for transformation operations.
package domain

// Transformer defines the interface for data transformation operations.
type Transformer interface {
	// Transform processes the input bytes and returns transformed bytes or an error.
	Transform(input []byte) ([]byte, error)
}
