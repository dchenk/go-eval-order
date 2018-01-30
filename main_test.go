package buf

import "testing"

func TestBuf(t *testing.T) {

	b := new(Buf)
	b.WriteByte('A')
	if b.Byte() != 'A' {
		t.Errorf("got wrong byte %s, not A", string(b.Byte()))
	}

}

