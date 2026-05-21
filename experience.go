package jsonresume

import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// Experience implements an embedded item in the JSON Resume "work" and "volunteer" fields arrays.
//
// In a JSON Resume document this might look like:
//
//	"work": [
//		{
//			"name": "SuperCo",
//			"url": "http://super.example",
//			"position": "Chief Technology Officer (CTO)",
//			"startDate": "2024-01-01",
//			"summary": "Technical leadership for a company of 50+ people.",
//			"highlights": [
//				"Hired initial team.",
//				"Created architecture.",
//				"Set up management structure.",
//			]
//		},
//		{
//			"name": "Acme",
//			"url": "http://acme.example",
//			"position": "Software Engineer",
//			"startDate": "2019-03-13",
//			"endDate": "2023-12-31",
//			"summary": "Create an electronic wallet.",
//			"highlights": [
//				"Built the back-end in Golang from scratch",
//				"Created payment system"
//			]
//		}
//	],
//	"volunteer": [
//		{
//			"organization": "Tech Conf",
//			"url": "http://techconf.example",
//			"position": "Organizer",
//			"startDate": "2025-02-01",
//			"endDate": "2026-08-24",
//			"summary": "Founded and organized Tech Conf.",
//			"highlights": [
//				"Organized Tech Conf",
//				"Brought in sponsors"
//			]
//		},
//		{
//			"organization": "Code School 123",
//			"url": "http://school123.example",
//			"position": "Instructor",
//			"startDate": "2018-03-13",
//			"endDate": "2024-07-07",
//			"summary": "Taught class on Golang.",
//			"highlights": [
//				"Taught 531 students the Go programming-language over a number of years.",
//			]
//		}
//	],
//
// Experience is an implementation of the individual JSON objects in that JSON array.
//
// Example usage:
//
//	var cv jsonresume.Resume
//	
//	// ...
//	
//	cv.Work = append(cv.Work, jsonresume.Experience{
//		Name:         nul.Something("SuperCo"),
//		URL:          nul.Something("http://super.example"),
//		Position:     nul.Something("Chief Technology Officer (CTO)"),
//		StartDate:    nul.Something("2024-01-01"),
//		Summary:      nul.Something(""Technical leadership for a company of 50+ people.),
//		Highlights:   activitypub.SomeString(
//			"Hired initial team.",
//			"Created architecture.",
//			"Set up management structure.",
//		),
//	})
//	
//	cv.Work = append(cv.Work, jsonresume.Experience{
//		Name:         nul.Something("Acme"),
//		URL:          nul.Something("http://acme.example"),
//		Position:     nul.Something("Software Engineer"),
//		StartDate:    nul.Something("2019-03-13"),
//		EndDate:      nul.Something("2023-12-31"),
//		Summary:      nul.Something("Create an electronic wallet."),
//		Highlights:   activitypub.SomeString(
//			"Built the back-end in Golang from scratch",
//			"Created payment system",
//		),
//	})
//	
//	// ...
//	
//	cv.Volunteer = append(cv.Volunteer, jsonresume.Experience{
//		Organization: nul.Something("Tech Conf"),
//		URL:          nul.Something("http://techconf.example"),
//		Position:     nul.Something("Organizer"),
//		StartDate:    nul.Something("2025-02-01"),
//		EndDate:      nul.Something("2026-08-24"),
//		Summary:      nul.Something("Founded and organized Tech Conf."),
//		Highlights:   activitypub.SomeString(
//			"Organized Tech Conf",
//			"Brought in sponsors",
//		),
//	})
//	
//	cv.Volunteer = append(cv.Volunteer, jsonresume.Experience{
//		Organization: nul.Something("Code School 123"),
//		URL:          nul.Something("http://school123.example"),
//		Position:     nul.Something("Instructor"),
//		StartDate:    nul.Something("2018-03-13"),
//		EndDate:      nul.Something("2024-07-07"),
//		Summary:      nul.Something("Taught class on Golang."),
//		Highlights:   activitypub.SomeString(
//			"Taught 531 students the Go programming-language over a number of years.",
//		),
//	})
type Experience struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/fep/6158"`
	Prefix    jsonld.Prefix    `jsonld:"cv"`

	ID   jsonld.ID          `json:"id,omitempty"`
	Type json.Const[string] `json:"type" json.value:"Experience"`

	activitypub.CoreEntity
	activitypub.CoreObject
	CoreExperience
}

func (receiver Experience) ProtoNode() activitypub.AnyNode {
	const _type string = TypeExperience

	return activitypub.AnyNode{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),
	}
}

func (receiver Experience) ProtoEntity() activitypub.AnyEntity {
	const _type string = TypeExperience

	return activitypub.AnyEntity{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity: receiver.CoreEntity,
	}
}

func (receiver Experience) ProtoObject() activitypub.AnyObject {
	const _type string = TypeExperience

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

func (receiver Experience) ProtoExperience() AnyExperience {
	const _type string = TypeExperience

	var result = AnyExperience{
		ID:   receiver.ID,
		Type: jsonld.SomeType(_type),

		CoreEntity:     receiver.CoreEntity,
		CoreObject:     receiver.CoreObject,
		CoreExperience: receiver.CoreExperience,
	}

	result.Attachments = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Attachments...)
	result.Tags = append([]activitypub.ProtoObjectOrProtoLink(nil), receiver.Tags...)

	return result
}
