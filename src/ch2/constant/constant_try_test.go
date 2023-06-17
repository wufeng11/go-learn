package constant

import "testing"

const (
	Mon = iota + 1
	Tue
	Wen
)

const (
	Readable = 1 << iota
	Writeable
	Execute
)

func TestConstantTry(t *testing.T) {
	t.Log(Mon, Tue, Wen)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, Writeable, Execute)
}
