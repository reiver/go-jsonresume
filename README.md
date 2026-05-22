# go-jsonresume

Package **jsonresume** provides tools for working with **JSON Resume** including **JSON Resume** expressed as JSON-LD (i.e., **FEP-6158**), for the Go programming language.

## Documentation

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-jsonresume

[![GoDoc](https://godoc.org/github.com/reiver/go-jsonresume?status.svg)](https://godoc.org/github.com/reiver/go-jsonresume)

## Example

To point to a resume from an ActivityPub actor, do something similar to the following:

```golang
import (
	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-jsonresume"
)

// ...

var actor activitypub.Person
//@TODO: set the values of the actor

var jsonResume json.JSONResume
jsonResume.SetResumeID("https://example.com/path/to/resume")

err := jsonld.Marshal(actor, jsonResume)
```

This will create ActivityPub JSON-LD similar to:

```json
{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "https://w3id.org/fep/6158",

    ...
  ],


  ...

  "type": "Person",

  ...

  "resume": "https://example.com/path/to/resume",

  ...
}
```

Notice 2 things in this ActivityPub JSON-LD.

№1: The new JSON-LD namespace: `https://w3id.org/fep/6158`

№2: The new field, `resume`, that whose value is the same value we set in the Go code (`https://example.com/path/to/resume`).

## Import

To import package **jsonresume** use `import` code like the following:
```
import "github.com/reiver/go-jsonresume"
```

## Installation

To install package **jsonresume** do the following:
```
GOPROXY=direct go get github.com/reiver/go-jsonresume
```

## Author

Package **jsonresume** was written by [Charles Iliya Krempeaux](http://reiver.link)

## See also:

* [JSON Resume](https://jsonresume.org/)
* [JSON Resume schema.json](https://github.com/jsonresume/resume-schema/blob/master/schema.json)
