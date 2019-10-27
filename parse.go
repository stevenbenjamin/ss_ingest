package ss_ingest

import "errors"

import (
	"regexp"
	"strings"
)

type Alert struct {
	Values  map[string]string
	Headers map[string]string
}

type Location struct {
	Street           string
	City             string
	State            string
	Zip              string
	FormattedAddress string
	Latitude         float64
	Longitude        float64
	XS               string
	Map              string
	Cad              string
	Grid             string
}

var (
	//parse delimiter lines in multipart form submission:  usually ^--xYzZY but not guaranteed
	//(delimiter is normally specified in multipart form header that we're not getting)
	delimiterRegex = regexp.MustCompile("(?m)^--.*$")

	//find string of the form "content-disposition"... "name=".." return stripped value of remainder e.g.
	//Content-Disposition: form-data; name="dkim"
	//{@emergycare.onmicrosoft.com : pass}
	sectionRegex = regexp.MustCompile("(?si)Content-Disposition.*name=\"([^\"]*)\"(.*)")

	//block of headers in the form "X-Forefront-PRVS: 0162ACCC24"
	headerRegex = regexp.MustCompile("(?m)^([^:]*):(.*)$")
)

//parse form field as a (lowercased-key) name value pair.
func parseSection(i string) (string, string, error) {
	res := sectionRegex.FindStringSubmatch(i)
	if len(res) < 3 { // no match
		return "", "", errors.New("no match")
	}
	return strings.ToLower(res[1]), strings.TrimSpace(res[2]), nil
}

//parse individual headers within the header form field
func parseHeaders(s string) map[string]string {
	headers := make(map[string]string)
	out := headerRegex.FindAllStringSubmatch(s, -1)
	for _, v := range out {
		headers[strings.ToLower(v[1])] = strings.TrimSpace(v[2])
	}
	return headers
}

//parse form-encoded alert email parsed into map structure
func Parse(i string) *Alert {
	a := &Alert{Values: make(map[string]string)}
	sections := delimiterRegex.Split(i, -1)
	for _, section := range sections {
		if k, v, err := parseSection(section); err == nil {
			a.Values[k] = v
		}
	}

	if headers, ok := a.Values["headers"]; ok {
		delete(a.Values, "headers")
		a.Headers = parseHeaders(headers)
	}

	return a
}
