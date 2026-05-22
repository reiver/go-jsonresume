package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CorePublication struct {
	Publisher   nul.Nullable[string] `json:"publisher"`
	ReleaseDate nul.Nullable[string] `json:"releaseDate"`
}
