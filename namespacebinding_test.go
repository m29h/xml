// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"testing"
)

func TestAdd(t *testing.T) {
	url := "http://binding-test-example.com/abc"
	err := NameSpaceBinding.Add(url, "xyz")
	if err != nil {
		t.Error(err)
	}

	type Person struct {
		XMLName Name `xml:"http://binding-test-example.com/abc person"`
		Id      int  `xml:"id,attr"`
	}

	v := &Person{Id: 13}
	want := `<xyz:person xmlns:xyz="http://binding-test-example.com/abc" id="13"></xyz:person>`
	got, err := Marshal(v)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got `%s`, want `%s`", got, want)
	}

}

func TestClear(t *testing.T) {
	url := "http://binding-test-example.com/abc"
	err := NameSpaceBinding.Add(url, "xyz")
	if err != nil {
		t.Error(err)
	}
	if NameSpaceBinding.get(url) != "xyz" {
		t.Error("binding was not set")
	}
	NameSpaceBinding.Clear()
	if NameSpaceBinding.get(url) == "xyz" {
		t.Error("binding was not cleared")
	}
}
