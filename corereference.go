package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreReference struct {
	Reference nul.Nullable[string] `json:"reference"`
}
