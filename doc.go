/*
Package jsonresume provides tools for working with JSON Resume (https://jsonresume.org/) including JSON Resume expressed as JSON-LD via FEP-6158 (https://w3id.org/fep/6158), for the Go programming language.

# Type System

Each JSON Resume entity (Award, Basics, Certificate, Education, Experience, Interest, Language, Location, Meta, Profile, Project, Publication, Reference, Resume, Skill) is represented by four types:

	• Main types (e.g., [Award], [Resume]) — for marshaling with a fixed JSON-LD type string.
	  Can also be used for unmarshaling when strict type validation is desired.
	• Core types (e.g., [CoreAward], [CoreResume]) — shared field definitions,
	  embedded by both main and Any types.
	• Any types (e.g., [AnyAward], [AnyResume]) — for unmarshaling with flexible
	  type acceptance. Use these when you do not need to validate the type value.
	• ID types (e.g., [AwardID], [ResumeID]) — JSON-LD references by IRI string.

# JSON-LD Integration

Types in this package carry JSON-LD namespace and prefix metadata via struct tags, enabling proper @context generation when marshaled with [github.com/reiver/go-jsonld].
Fields shared with ActivityPub/ActivityStreams (name, summary, url, image) are mapped to the "as:" prefix, while JSON Resume fields use the "cv:" prefix.

# Attaching a Resume to an ActivityPub Actor

Use [JSONResume] to attach resume references to an ActivityPub actor:

	var jsonResume jsonresume.JSONResume
	jsonResume.SetResumeID("https://example.com/path/to/resume")
	
	bytes, err := jsonld.Marshal(actor, jsonResume)

# Design Decisions

Project uses "@type" (not "type") for the ActivityPub/ActivityStreams JSON-LD type field.
This is because JSON Resume defines a "type" field on projects for the project category (e.g., "application", "library"), which collides with JSON-LD's use of "type" as an alias for "@type".
To preserve both, [Project] and [AnyProject] use "@type" for the JSON-LD type, while [CoreProject].ProjectType holds the project category.
All other types use "type" for the ActivityPub/ActivityStreams JSON-LD type.
*/
package jsonresume
