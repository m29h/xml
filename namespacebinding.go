package xml

import (
	"fmt"
	"strings"
)

var NameSpaceBinding hint

type hint struct {
	nsPrefixHint map[string]string
}

func init() {
	NameSpaceBinding.nsPrefixHint = map[string]string{
		"http://www.w3.org/2001/XMLSchema":                                                   "xsd",
		"http://www.w3.org/2001/XMLSchema-instance":                                          "xsi",
		"http://schemas.xmlsoap.org/soap/envelope/":                                          "soapenv",
		"http://www.w3.org/2000/09/xmldsig#":                                                 "ds",
		"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd":  "wsse",
		"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd": "wsu",
		"http://www.w3.org/2001/10/xml-exc-c14n#":                                            "ec",
	}
}

// clears all stored namespace bindings including the opinionated default binding
func (h *hint) Clear() {
	NameSpaceBinding.nsPrefixHint = make(map[string]string)
}

// Forces binding a namespace url to a prefix
func (h *hint) Add(url string, prefix string) error {
	prefix = strings.TrimSpace(prefix)
	// ignore/refuse an illegal prefix
	if prefix == "" || !isName([]byte(prefix)) || strings.Contains(prefix, ":") {
		return fmt.Errorf("prefix `%s` contains illegal sequence", prefix)
	}
	// xmlanything is reserved and any variant of it regardless of
	// case should be matched, so:
	//    (('X'|'x') ('M'|'m') ('L'|'l'))
	// See Section 2.3 of https://www.w3.org/TR/REC-xml/
	if len(prefix) >= 3 && strings.EqualFold(prefix[:3], "xml") {
		return fmt.Errorf("prefix `%s` may not start with xml..", prefix)
	}

	h.nsPrefixHint[strings.TrimSpace(url)] = prefix
	return nil
}

func (h *hint) get(url string) string {
	if prefix, ok := h.nsPrefixHint[url]; ok {
		return prefix
	}

	// Pick a name. We try to use the final element of the path
	// but fall back to _.
	prefix := strings.TrimRight(url, "/")
	if i := strings.LastIndex(prefix, "/"); i >= 0 {
		prefix = prefix[i+1:]
	}
	if prefix == "" || !isName([]byte(prefix)) || strings.Contains(prefix, ":") {
		prefix = "_"
	}
	// xmlanything is reserved and any variant of it regardless of
	// case should be matched, so:
	//    (('X'|'x') ('M'|'m') ('L'|'l'))
	// See Section 2.3 of https://www.w3.org/TR/REC-xml/
	if len(prefix) >= 3 && strings.EqualFold(prefix[:3], "xml") {
		prefix = "_" + prefix
	}
	return prefix
}
