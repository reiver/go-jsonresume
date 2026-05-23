package jsonresume

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrBytesEmpty      = erorr.Error("empty bytes")
	ErrTypeUnsupported = erorr.Error("unsupported type")
)
