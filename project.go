package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Project implements an embedded item in the JSON Resume "projects" fields array.
//
// In a JSON Resume document this might look like:
//
//	"projects": [
//		{
//			"name": "Microdon",
//			"url": "https://codeberg.org/reiver/microdon"
//			"startDate": "2024-11-16",
//			"description": "lightweight ActivityPub back-end",
//			"highlights": [
//				"created ActivityPub back-end from scratch"
//			]
//		},
//		{
//			"name": "SpaceMax",
//			"url": "https://spacemax.example"
//			"startDate": "2011-03-02",
//			"endDate": "2015-08-17",
//			"description": "virtual drive",
//			"highlights": [
//				"reached more-than 100,000 users",
//				"released version 2"
//			]
//		}
//	],
//
// Project is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Projects = append(cv.Projects, jsonresume.Project{
//		Name:        nul.Something("Microdon"),
//		URL:         activitypub.HRef("https://codeberg.org/reiver/microdon"),
//		StartDate:   nul.Something("2024-11-16"),
//		Description: nul.Something("lightweight ActivityPub back-end"),
//		Highlights:  activitypub.SomeStrings(
//			"created ActivityPub back-end from scratch"
//		),
//	})
//	
//	cv.Projects = append(cv.Projects, jsonresume.Project{
//		Name:        nul.Something("SpaceMax"),
//		URL:         activitypub.HRef("https://spacemax.example"),
//		StartDate:   nul.Something("2011-03-02"),
//		EndDate:     nul.Something("2015-08-17"),
//		Description: nul.Something("virtual drive"),
//		Highlights:  activitypub.SomeStrings(
//			"reached more-than 100,000 users",
//			"released version 2"
//		),
//	})
type Project struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Project"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreProject
}

func (receiver Project) ProtoNode() activitypub.AnyNode {
	const _type string = TypeProject

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Project) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeProject

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Project) ProtoObject() activitypub.AnyObject {
	const _type string = TypeProject

	var result = activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}

func (receiver Project) ProtoProject() AnyProject {
	const _type string = TypeProject

	var result = AnyProject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
		CoreObject: receiver.CoreObject,
		CoreProject:  receiver.CoreProject,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
