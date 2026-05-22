package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

type CoreBasics struct {
	EMail    activitypub.Strings  `json:"email"`
	Label    activitypub.Strings `json:"label"`
	Phone    activitypub.Strings `json:"phone"`
	Profiles []ProtoProfile      `json:"profiles"`
}
