// package circular implements routines of a circular buffer of bytes
// supporting both overflow-checked writes and unconditional, possibly
// overwriting, writes.
//
// The provided API is chosen so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

package circular

import "errors"

type Buffer struct {
	data   []byte
	size   int
	pos    int
	next   int
	stored int
}

// NewBuffer creates a new buffer
func NewBuffer(size int) *Buffer {
	if size <= 0 {
		return &Buffer{data: []byte{}}
	}
	return &Buffer{data: make([]byte, size), size: size}
}

// ReadByte reads off a byte from the buffer
func (b *Buffer) ReadByte() (byte, error) {
	if (*b).stored == 0 {
		return byte(0), errors.New("none stored")
	}
	data := (*b).data[(*b).pos]
	(*b).stored--
	(*b).pos = ((*b).pos + 1) % ((*b).size)
	return data, nil
}

// WriteByte write byte value in the buffer
func (b *Buffer) WriteByte(c byte) error {
	if (*b).stored < (*b).size {
		(*b).stored++
		(*b).data[(*b).next] = c
		(*b).next = ((*b).next + 1) % (*b).size
		return nil
	}
	return errors.New("buffer is full")
}

// Overwrite writes byte in the buffer even if it is full.
// The earliest byte value written is overwritten
func (b *Buffer) Overwrite(c byte) {
	if (*b).stored < (*b).size {
		(*b).stored++
	} else {
		(*b).pos = ((*b).pos + 1) % (*b).size
	}
	(*b).data[(*b).next] = c
	(*b).next = ((*b).next + 1) % (*b).size
}

// Reset resets the buffer to initial state
func (b *Buffer) Reset() {
	(*b).stored, (*b).pos, (*b).next = 0, 0, 0
}
