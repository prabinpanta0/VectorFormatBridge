package svg

import (
	"encoding/xml"
	"io/ioutil"
)

// ParseSVG reads and parses an SVG file
func ParseSVG(filename string) (*SVG, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var svg SVG
	err = xml.Unmarshal(data, &svg)
	if err != nil {
		return nil, err
	}

	return &svg, nil
}

// WriteSVG writes SVG content to a file
func WriteSVG(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
