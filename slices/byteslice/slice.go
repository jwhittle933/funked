package byteslice

import "bytes"

type Slice []byte

func From(in []byte) Slice {
	return in
}

func (s Slice) Compare(compareTo []byte) int {
	return bytes.Compare(s, compareTo)
}

func (s Slice) Contains(subslice []byte) bool {
	return bytes.Contains(s, subslice)
}

func (s Slice) ContainsAny(chars string) bool {
	return bytes.ContainsAny(s, chars)
}

func (s Slice) ContainsRune(r rune) bool {
	return bytes.ContainsRune(s, r)
}

func (s Slice) Count(sep []byte) int {
	return bytes.Count(s, sep)
}

func (s Slice) Equal(b []byte) bool {
	return bytes.Equal(s, b)
}

func (s Slice) EqualFolds(t []byte) bool {
	return bytes.EqualFold(s, t)
}

func (s Slice) Bytes() []byte {
	return s
}
