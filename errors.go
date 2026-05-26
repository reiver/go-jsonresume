package jsonresume

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrBytesEmpty      = erorr.Error("empty bytes")
	ErrReceiverNil     = erorr.Error("nil receiver")
	ErrTypeUnsupported = erorr.Error("unsupported type")
)
