package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreAward struct {
	Awarder nul.Nullable[string] `json:"awarder"`
	Date    nul.Nullable[string] `json:"date"`
	Summary nul.Nullable[string] `json:"summary"`
	Title   nul.Nullable[string] `json:"title"`
}
