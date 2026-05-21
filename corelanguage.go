package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreLanguage struct {
	Fluency  nul.Nullable[string] `json:"fluency"`
	Language nul.Nullable[string] `json:"language"`
}
