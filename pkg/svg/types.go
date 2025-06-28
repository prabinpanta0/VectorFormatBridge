package svg

import "encoding/xml"

// SVG represents the root SVG element and its child elements
type SVG struct {
	XMLName   xml.Name   `xml:"svg"`
	Width     string     `xml:"width,attr"`
	Height    string     `xml:"height,attr"`
	Rects     []Rect     `xml:"rect"`
	Circles   []Circle   `xml:"circle"`
	Lines     []Line     `xml:"line"`
	Paths     []Path     `xml:"path"`
	Ellipses  []Ellipse  `xml:"ellipse"`
	Polygons  []Polygon  `xml:"polygon"`
	Polylines []Polyline `xml:"polyline"`
}

// Rect represents an SVG rectangle element
type Rect struct {
	X      string `xml:"x,attr"`
	Y      string `xml:"y,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Fill   string `xml:"fill,attr"`
	Stroke string `xml:"stroke,attr"`
}

// Circle represents an SVG circle element
type Circle struct {
	Cx     string `xml:"cx,attr"`
	Cy     string `xml:"cy,attr"`
	R      string `xml:"r,attr"`
	Fill   string `xml:"fill,attr"`
	Stroke string `xml:"stroke,attr"`
}

// Line represents an SVG line element
type Line struct {
	X1     string `xml:"x1,attr"`
	Y1     string `xml:"y1,attr"`
	X2     string `xml:"x2,attr"`
	Y2     string `xml:"y2,attr"`
	Stroke string `xml:"stroke,attr"`
}

// Path represents an SVG path element
type Path struct {
	D      string `xml:"d,attr"`
	Stroke string `xml:"stroke,attr"`
	Fill   string `xml:"fill,attr"`
}

// Ellipse represents an SVG ellipse element
type Ellipse struct {
	Cx     string `xml:"cx,attr"`
	Cy     string `xml:"cy,attr"`
	Rx     string `xml:"rx,attr"`
	Ry     string `xml:"ry,attr"`
	Fill   string `xml:"fill,attr"`
	Stroke string `xml:"stroke,attr"`
}

// Polygon represents an SVG polygon element
type Polygon struct {
	Points string `xml:"points,attr"`
	Fill   string `xml:"fill,attr"`
	Stroke string `xml:"stroke,attr"`
}

// Polyline represents an SVG polyline element
type Polyline struct {
	Points string `xml:"points,attr"`
	Stroke string `xml:"stroke,attr"`
}
