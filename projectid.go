package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
)

// ProjectID implements a referenced item in the JSON Resume "projects" fields array.
//
// This is something that one could seen in a JSON-LD expression of JSON Resume, rather than plain JSON Resume.
//
// In a JSON-LD flavored JSON Resume document, this might look like:
//
//	"projects": [
//		"http://example.com/resume/project/microdon",
//		"http://example.com/resume/project/spacemax",
//	],
//
// ProjectID is an implementation of the individual JSON strings in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Projects = append(cv.Projects, jsonresume.ProjectID("http://example.com/resume/project/microdon"))
//	
//	cv.Projects = append(cv.Projects, jsonresume.ProjectID("http://example.com/resume/project/spacemax"))
type ProjectID string

func (receiver ProjectID) ProtoNode() activitypub.AnyNode {
	const _type string = TypeProject

	return activitypub.AnyNode{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ProjectID) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeProject

	return activitypub.AnyEntity{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ProjectID) ProtoObject() activitypub.AnyObject {
	const _type string = TypeProject

	return activitypub.AnyObject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

func (receiver ProjectID) ProtoProject() AnyProject {
	const _type string = TypeProject

	return AnyProject{
		ID:   jsonld.SomeID(string(receiver)),
		Type: jsonld.SomeType(_type),
	}
}

