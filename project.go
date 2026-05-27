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
// Note that you should use Project for marshaling but not unmarshaling.
// For unmarshaling instead use [AnyProject].
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
		var bb []byte = []byte(raw.ID)

		if 0 < len(bb) {
			err := jsonld.Unmarshal(bb, &receiver.ID)
			if nil != err {
				err = erorr.Wrap(err, "failed to json-unmarshal project id")
				return err
			}
		}
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

	{
		{
			var bb []byte = []byte(raw.Description)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Description)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project description")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.EndDate)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.EndDate)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project endDate")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Entity)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Entity)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project entity")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Highlights)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Highlights)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project highlights")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Keywords)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Keywords)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project keywords")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Name)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Name)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project name")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.Roles)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.Roles)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project roles")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.StartDate)

			if 0 < len(bb) {
				err := jsonld.Unmarshal(bb, &receiver.StartDate)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project startDate")
					return err
				}
			}
		}

		{
			var bb []byte = []byte(raw.URL)

			if 0 < len(bb) {
				err := jsonld.UnmarshalJSONStringOrJSONObjectOrJSONArray[activitypub.ProtoLink, activitypub.HRef, activitypub.AnyLink](bb, &receiver.URL)
				if nil != err {
					err = erorr.Wrap(err, "failed to json-unmarshal project url")
					return err
				}
			}
		}
	}

	return nil
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
