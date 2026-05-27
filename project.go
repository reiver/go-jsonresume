package jsonresume

import (
	gojson "encoding/json"

	"codeberg.org/reiver/go-activitypub"
	"codeberg.org/reiver/go-erorr"
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
//			"created ActivityPub back-end from scratch",
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
//			"released version 2",
//		),
//	})
//
// Project is for marshaling with a fixed type of "Project".
// It can also be used for unmarshaling when strict type validation is desired —
// it rejects any type value other than "Project", "cv:Project", or "https://w3id.org/fep/6158#Project".
//
// For unmarshaling that accepts any type value, use [AnyProject] instead.
//
// See also:
//
//	• [AnyProject]
//	• [CoreProject]
//	• [ProjectID]
//	• [ProtoProject]
//	• [TypeProject]
type Project struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Project"`

	CoreProject
}

func (receiver *Project) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrReceiverNil
	}

	var raw rawProject
	err := jsonld.Unmarshal(bytes, &raw)
	if nil != err {
		err = erorr.Wrap(err, "failed to json-unmarshal project")
		return err
	}

	{
		var bb []byte = []byte(raw.Type)

		if 0 < len(bb) {
			var typeValue string
			err := gojson.Unmarshal(bb, &typeValue)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project type")
				return err
			}

			switch typeValue {
			case TypeProject, CompactTypeProject, ExpandedTypeProject:
				// OK
			default:
				return erorr.Errorf("jsonresume: unexpected type for project: %q", typeValue)
			}
		}
	}

	return receiver.CoreProject.unmarshalRawProject(raw, &receiver.ID)
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

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
	}
}

func (receiver Project) ProtoObject() activitypub.AnyObject {
	const _type string = TypeProject

	return activitypub.AnyObject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: activitypub.CoreEntity{
			Name: receiver.Name,
		},
		CoreObject: activitypub.CoreObject{
			URL: receiver.URL,
		},
	}
}

func (receiver Project) ProtoProject() AnyProject {
	const _type string = TypeProject

	return AnyProject{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreProject: receiver.CoreProject,
	}
}
