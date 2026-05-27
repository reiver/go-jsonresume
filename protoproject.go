package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
)

// ProtoProject represents something that is any type 'Project', including sub-types.
//
// See also:
//
//	• [AnyProject]
//	• [Project]
//	• [ProjectID]
type ProtoProject interface {
	activitypub.ProtoObject
	ProtoProject() AnyProject
}

var (
	_ ProtoProject = AnyProject{}
	_ ProtoProject = Project{}
	_ ProtoProject = ProjectID{}
)
