package jsonresume_test

import (
	"fmt"

	"github.com/reiver/go-jsonld"

	"github.com/reiver/go-jsonresume"
)

func ExampleAnyResume_jsonUnmarshal() {
	const jsonString string =
`{
  "id":"http://example.com/resume/1",
  "type":"Resume",
  "awards":[
    {
      "title":"Best Employee (2024)",
      "date":"2024-05-21",
      "awarder":"SuperCo",
      "summary":"He did good work."
    },
    {
      "title":"Acme Excellence (2021)",
      "date":"2021-01-17",
      "awarder":"Acme",
      "summary":"For Joe Blow's excellent work."
    }
  ],
  "certificates":[
    {
      "name": "CP+",
      "date": "2018-06-01",
      "issuer": "CPEdu",
      "url": "http://cp.example"
    },
    {
      "name": "Event Sourcing Core",
      "date": "2017-03-12",
      "issuer": "EngWorks",
      "url": "http://engworks.example"
    }
  ],
  "references":[
    {
      "name": "Jane Doe",
      "reference": "Joe Blow is an exceptionally talented professional."
    },
    {
      "name": "Bob Smith",
      "reference": "I enjoyed working with Joe Blow and give them my highest recommendation without reservation."
    }
  ],
  "skills":[
    {
      "name": "Photography",
      "keywords": ["Astrophotography", "Food", "Nature"]
    },
    {
      "name": "Woodworking",
      "keywords": ["Furniture"]
    }
  ],
  "work":[
    {
      "name": "SuperCo",
      "position": "Chief Technology Officer (CTO)",
      "startDate": "2024-01-01",
      "summary": "Technical leadership for a company of 50+ people.",
      "highlights": [
        "Hired initial team.",
        "Created architecture.",
        "Set up management structure."
      ]
    },
    {
      "name": "Acme",
      "position": "Software Engineer",
      "startDate": "2019-03-13",
      "endDate": "2023-12-31",
      "summary": "Create an electronic wallet.",
      "highlights": [
        "Built the back-end in Golang from scratch",
        "Created payment system"
      ]
    }
  ]
}
`
	var resume jsonresume.AnyResume

	err := jsonld.Unmarshal([]byte(jsonString), &resume)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Println("types:", resume.Type.Strings())
	fmt.Println("number of awards:", len(resume.Awards))
	fmt.Println("number of certificates:", len(resume.Certificates))
	fmt.Println("number of education:", len(resume.Education))
	fmt.Println("number of interests:", len(resume.Interests))
	fmt.Println("number of languages:", len(resume.Languages))
	fmt.Println("number of projects:", len(resume.Projects))
	fmt.Println("number of publications:", len(resume.Publications))
	fmt.Println("number of references:", len(resume.References))
	fmt.Println("number of skills:", len(resume.Skills))
	fmt.Println("number of volunteer:", len(resume.Volunteer))
	fmt.Println("number of work:", len(resume.Work))
	fmt.Println()
	fmt.Println(resume)

	// Output:
	// types: [Resume]
	// number of awards: 2
	// number of certificates: 2
	// number of education: 0
	// number of interests: 0
	// number of languages: 0
	// number of projects: 0
	// number of publications: 0
	// number of references: 2
	// number of skills: 2
	// number of volunteer: 0
	// number of work: 2
	//
	// {
	//   "@context": {
	//     "cv": "https://w3id.org/fep/6158",
	//     "awards": "cv:awards",
	//     "basics": "cv:basics",
	//     "certificates": "cv:certificates",
	//     "education": "cv:education",
	//     "id": "cv:id",
	//     "interests": "cv:interests",
	//     "languages": "cv:languages",
	//     "meta": "cv:meta",
	//     "projects": "cv:projects",
	//     "publications": "cv:publications",
	//     "references": "cv:references",
	//     "skills": "cv:skills",
	//     "type": "cv:type",
	//     "volunteer": "cv:volunteer",
	//     "work": "cv:work"
	//   },
	//   "id": "http://example.com/resume/1",
	//   "type": "Resume",
	//   "awards": [
	//     {
	//       "awarder": "SuperCo",
	//       "date": "2024-05-21",
	//       "summary": "He did good work.",
	//       "title": "Best Employee (2024)"
	//     },
	//     {
	//       "awarder": "Acme",
	//       "date": "2021-01-17",
	//       "summary": "For Joe Blow's excellent work.",
	//       "title": "Acme Excellence (2021)"
	//     }
	//   ],
	//   "certificates": [
	//     {
	//       "date": "2018-06-01",
	//       "name": "CP+",
	//       "issuer": "CPEdu",
	//       "url": "http://cp.example"
	//     },
	//     {
	//       "date": "2017-03-12",
	//       "name": "Event Sourcing Core",
	//       "issuer": "EngWorks",
	//       "url": "http://engworks.example"
	//     }
	//   ],
	//   "references": [
	//     {
	//       "name": "Jane Doe",
	//       "reference": "Joe Blow is an exceptionally talented professional."
	//     },
	//     {
	//       "name": "Bob Smith",
	//       "reference": "I enjoyed working with Joe Blow and give them my highest recommendation without reservation."
	//     }
	//   ],
	//   "skills": [
	//     {
	//       "keywords": [
	//         "Astrophotography",
	//         "Food",
	//         "Nature"
	//       ],
	//       "name": "Photography"
	//     },
	//     {
	//       "keywords": "Furniture",
	//       "name": "Woodworking"
	//     }
	//   ],
	//   "work": [
	//     {
	//       "highlights": [
	//         "Hired initial team.",
	//         "Created architecture.",
	//         "Set up management structure."
	//       ],
	//       "name": "SuperCo",
	//       "position": "Chief Technology Officer (CTO)",
	//       "startDate": "2024-01-01",
	//       "summary": "Technical leadership for a company of 50+ people."
	//     },
	//     {
	//       "endDate": "2023-12-31",
	//       "highlights": [
	//         "Built the back-end in Golang from scratch",
	//         "Created payment system"
	//       ],
	//       "name": "Acme",
	//       "position": "Software Engineer",
	//       "startDate": "2019-03-13",
	//       "summary": "Create an electronic wallet."
	//     }
	//   ]
	// }
}
