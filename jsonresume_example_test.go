package jsonresume_test

import (
	"fmt"

	"github.com/reiver/go-json"

	"github.com/reiver/go-jsonresume"
)

func ExampleJSONResume() {

	var obj jsonresume.JSONResume

	obj.Resume = append(obj.Resume, jsonresume.SomeResumeID("http://example.com/resume/executive"))
	obj.Resume = append(obj.Resume, jsonresume.SomeResumeID("http://example.com/resume/programmer"))

	bytes, err := json.Marshal(obj)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(bytes))

	// Output:
	// {"resume":["http://example.com/resume/executive","http://example.com/resume/programmer"]}
}
