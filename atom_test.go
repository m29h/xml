// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import "time"

var atomValue = &Feed{
	XMLName: Name{"http://www.w3.org/2005/Atom", "feed"},
	Title:   "Example Feed",
	Link:    []Link{{Href: "http://example.org/"}},
	Updated: ParseTime("2003-12-13T18:30:02Z"),
	Author:  Person{Name: "John Doe"},
	ID:      "urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6",

	Entry: []Entry{
		{
			Title:   "Atom-Powered Robots Run Amok",
			Link:    []Link{{Href: "http://example.org/2003/12/13/atom03"}},
			ID:      "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
			Updated: ParseTime("2003-12-13T18:30:02Z"),
			Summary: NewText("Some text."),
		},
	},
}

var atomXML = `<Atom:feed xmlns:Atom="http://www.w3.org/2005/Atom" updated="2003-12-13T18:30:02Z"><Atom:title>Example Feed</Atom:title><Atom:id>urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6</Atom:id><Atom:link href="http://example.org/"></Atom:link><Atom:author><Atom:name>John Doe</Atom:name><Atom:uri></Atom:uri><Atom:email></Atom:email></Atom:author><Atom:entry><Atom:title>Atom-Powered Robots Run Amok</Atom:title><Atom:id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</Atom:id><Atom:link href="http://example.org/2003/12/13/atom03"></Atom:link><Atom:updated>2003-12-13T18:30:02Z</Atom:updated><Atom:author><Atom:name></Atom:name><Atom:uri></Atom:uri><Atom:email></Atom:email></Atom:author><Atom:summary>Some text.</Atom:summary></Atom:entry></Atom:feed>`

func ParseTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	return t
}

func NewText(text string) Text {
	return Text{
		Body: text,
	}
}
