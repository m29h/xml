# xml [![GoDoc](https://godoc.org/github.com/m29h/xml?status.svg)](https://godoc.org/github.com/m29h/xml) [![Tests](https://github.com/m29h/xml/actions/workflows/test.yml/badge.svg)](https://github.com/m29h/xml/actions/workflows/test.yml) [![Coverage Status](https://coveralls.io/repos/github/m29h/xml/badge.svg)](https://coveralls.io/github/m29h/xml) [![Go Report Card](https://goreportcard.com/badge/github.com/m29h/xml)](https://goreportcard.com/report/github.com/m29h/xml)

Package `github.com/m29h/xml` implements a XML 1.0 marshaler that is a drop in `encoding/xml` replacement maintaining full interface compatibility to it but aims to producing [C14N Exclusive XML Canonicalization](https://www.w3.org/TR/xml-exc-c14n/) compatible byte sequences. This includes sorting of all rendered element attributes as per the C14N-XML spec.

## Example
A Struct defined in the normal way as known from `encoding/xml`

```go
type Person struct {
	XMLName xml.Name `xml:"http://example.com/ns1 person"`
	Name    string   `xml:"name"`
	Phone   string   `xml:"http://example.com/ns2 phone,omitempty"`
}
```
is marshaled to
```xml
<ns1:person xmlns:ns1="http://example.com/ns1">
  <ns1:name>Foo</ns1:name>
  <ns2:phone xmlns:ns2="http://example.com/ns2">123</ns2:phone>
</ns1:person>
```
