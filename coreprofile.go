package jsonresume

import (
	"github.com/reiver/go-nul"
)

type CoreProfile struct {
        Network  nul.Nullable[string] `json:"network"`
        UserName nul.Nullable[string] `json:"username"`
}
