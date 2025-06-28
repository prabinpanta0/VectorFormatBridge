package transform

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Transform represents a transformation matrix with translation, scale, and rotation
type Transform struct {
	X, Y, Scale, Rotate float64
}

// NewTransform creates a new transform with default values
func NewTransform() Transform {
	return Transform{0, 0, 1, 0}
}

// ApplyToPoint applies the transform to a point (x, y) and returns the transformed coordinates
func (t Transform) ApplyToPoint(x, y float64) (float64, float64) {
	// If scale is 0, set it to 1 to avoid issues
	scale := t.Scale
	if scale == 0 {
		scale = 1
	}

	// Scale
	x *= scale
	y *= scale

	// Rotate
	angle := t.Rotate * (math.Pi / 180)
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	xRot := x*cos - y*sin
	yRot := x*sin + y*cos

	// Translate
	xFinal := xRot + t.X
	yFinal := yRot + t.Y

	return xFinal, yFinal
}

// ParseTransform parses a transform string like T(x,y,scale,rotate)
func ParseTransform(s string) Transform {
	r := regexp.MustCompile(`T\(([^,]+),([^,]+),([^,]+),([^)]+)\)`)
	m := r.FindStringSubmatch(s)
	if len(m) != 5 {
		return NewTransform() // Default: no translation, scale=1, no rotation
	}
	x, _ := strconv.ParseFloat(strings.TrimSpace(m[1]), 64)
	y, _ := strconv.ParseFloat(strings.TrimSpace(m[2]), 64)
	scale, _ := strconv.ParseFloat(strings.TrimSpace(m[3]), 64)
	if scale == 0 {
		scale = 1 // Default scale
	}
	rot, _ := strconv.ParseFloat(strings.TrimSpace(m[4]), 64)
	return Transform{x, y, scale, rot}
}
