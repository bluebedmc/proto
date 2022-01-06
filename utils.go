package proto

import "io"

// readByte reads one byte from io.Reader
func readByte(r io.Reader) (byte, error) {
	if r, ok := r.(io.ByteReader); ok {
		return r.ReadByte()
	}
	var v [1]byte
	_, err := io.ReadFull(r, v[:])
	return v[0], err
}
