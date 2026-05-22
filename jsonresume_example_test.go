package jsonresume_test

import (
	"fmt"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonld"

	"github.com/reiver/go-jsonresume"
)

func ExampleJSONResume() {

	var jsonResume jsonresume.JSONResume

	jsonResume.AppendResumeIRI("http://example.com/resume/executive")
	jsonResume.AppendResumeIRI("http://example.com/resume/programmer")

	bytes, err := jsonld.Marshal(activitypub.SomeName("Joe Blow"), jsonResume)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"@context":{"as":"https://www.w3.org/ns/activitystreams","cv":"https://w3id.org/fep/6158","name":"as:name","resume":"cv:resume"},"name":"Joe Blow","resume":["http://example.com/resume/executive","http://example.com/resume/programmer"]}
}
