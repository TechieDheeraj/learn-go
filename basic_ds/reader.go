package main

import "bytes"

type MyReader struct{}

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(p []byte) (int, error) {
	a := []byte{'A'}
	copy(p, bytes.Repeat(a, len(p)))
	return len(p), nil
}
