package buf

// A Buf can be used to write one byte at a time to the underlying slice.
type Buf struct {
	data []byte
}

// WriteByte appends bt to the buffer.
func (b *Buf) WriteByte(bt byte) {
	b.data[b.grow()] = bt // IS THIS HERE GUARANTEED TO WORK AS EXPECTED? Must b.grow be called first?
}

func (b *Buf) grow() int {
	// This is a trivial example, but something like this is the bytes.Buffer.grow method,
	// which grows the buffer's underlying slice and returns the index at which to write.
	indx := len(b.data)
	b.data = append(b.data, 0)
	return indx
}

// Byte returns the first byte of the buffer or '-' if it's empty.
func (b *Buf) Byte() byte {
	if len(b.data) == 0 {
		return '-'
	}
	return b.data[0]
}
