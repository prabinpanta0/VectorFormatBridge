package converter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/prabinpanta0/VectorFormatBridge/pkg/transform"
)

// extractParams extracts parameters from a command line like "M(800,600,#fff)"
func extractParams(line string) []string {
	// Find content between parentheses
	start := strings.Index(line, "(")
	end := strings.Index(line, ")")
	if start == -1 || end == -1 || start >= end {
		return []string{}
	}

	content := line[start+1 : end]
	parts := strings.Split(content, ",")

	// Trim spaces from each part
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts
}

// parseF parses a string to float64
func parseF(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return val
}

// extractStyle extracts style information from a line
func extractStyle(line string) string {
	// Look for S(...) pattern
	re := regexp.MustCompile(`S\(([^)]+)\)`)
	match := re.FindStringSubmatch(line)
	if len(match) < 2 {
		return `stroke="black" fill="none"`
	}

	params := strings.Split(match[1], ",")
	style := ""

	if len(params) > 0 && strings.TrimSpace(params[0]) != "#none" {
		style += fmt.Sprintf(`stroke="%s"`, strings.TrimSpace(params[0]))
	}

	if len(params) > 1 && strings.TrimSpace(params[1]) != "#none" {
		if style != "" {
			style += " "
		}
		style += fmt.Sprintf(`fill="%s"`, strings.TrimSpace(params[1]))
	}

	if style == "" {
		style = `stroke="black" fill="none"`
	}

	return style
}

// extractPathData extracts path data from P[...] format
func extractPathData(line string) string {
	start := strings.Index(line, "[")
	end := strings.LastIndex(line, "]")
	if start == -1 || end == -1 || start >= end {
		return ""
	}
	return line[start+1 : end]
}

// extractPointList extracts point list from PG[...] or PL[...] format
func extractPointList(line string) string {
	start := strings.Index(line, "[")
	end := strings.LastIndex(line, "]")
	if start == -1 || end == -1 || start >= end {
		return ""
	}
	return line[start+1 : end]
}

// transformPoints applies transform to a list of points
func transformPoints(points string, t transform.Transform) string {
	if points == "" {
		return ""
	}

	coords := strings.Fields(strings.ReplaceAll(points, ",", " "))
	var transformed []string

	for i := 0; i < len(coords)-1; i += 2 {
		x := parseF(coords[i])
		y := parseF(coords[i+1])

		newX, newY := t.ApplyToPoint(x, y)
		transformed = append(transformed, fmt.Sprintf("%.2f,%.2f", newX, newY))
	}

	return strings.Join(transformed, " ")
}

// extractGroupContent extracts content from G[...] format
func extractGroupContent(line string) string {
	start := strings.Index(line, "[")
	end := strings.LastIndex(line, "]")
	if start == -1 || end == -1 || start >= end {
		return "<!-- Empty group -->"
	}
	return line[start+1 : end]
}

// renderEntity renders an entity with transform
func renderEntity(entity string, t transform.Transform) string {
	return renderLine(entity, t)
}

// renderLine renders a single EGF line as SVG
func renderLine(line string, t transform.Transform) string {
	switch {
	case strings.HasPrefix(line, "R("):
		p := extractParams(line)
		if len(p) < 4 {
			return fmt.Sprintf("<!-- Invalid rect: %s -->", line)
		}
		x, y := t.ApplyToPoint(parseF(p[0]), parseF(p[1]))
		w, h := parseF(p[2])*t.Scale, parseF(p[3])*t.Scale
		style := extractStyle(line)
		return fmt.Sprintf(`<rect x="%f" y="%f" width="%f" height="%f" %s/>`, x, y, w, h, style)

	case strings.HasPrefix(line, "C("):
		p := extractParams(line)
		if len(p) < 3 {
			return fmt.Sprintf("<!-- Invalid circle: %s -->", line)
		}
		x, y := t.ApplyToPoint(parseF(p[0]), parseF(p[1]))
		r := parseF(p[2]) * t.Scale
		style := extractStyle(line)
		return fmt.Sprintf(`<circle cx="%f" cy="%f" r="%f" %s/>`, x, y, r, style)

	case strings.HasPrefix(line, "L("):
		p := extractParams(line)
		if len(p) < 4 {
			return fmt.Sprintf("<!-- Invalid line: %s -->", line)
		}
		x1, y1 := t.ApplyToPoint(parseF(p[0]), parseF(p[1]))
		x2, y2 := t.ApplyToPoint(parseF(p[2]), parseF(p[3]))
		style := extractStyle(line)
		return fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" %s/>`, x1, y1, x2, y2, style)

	case strings.HasPrefix(line, "P["):
		path := extractPathData(line)
		style := extractStyle(line)
		// Path transform is skipped for now — would require parsing path commands
		return fmt.Sprintf(`<path d="%s" %s/>`, path, style)

	case strings.HasPrefix(line, "E("):
		p := extractParams(line)
		if len(p) < 4 {
			return fmt.Sprintf("<!-- Invalid ellipse: %s -->", line)
		}
		cx, cy := t.ApplyToPoint(parseF(p[0]), parseF(p[1]))
		rx := parseF(p[2]) * t.Scale
		ry := parseF(p[3]) * t.Scale
		style := extractStyle(line)
		return fmt.Sprintf(`<ellipse cx="%f" cy="%f" rx="%f" ry="%f" %s/>`, cx, cy, rx, ry, style)

	case strings.HasPrefix(line, "PG["):
		points := transformPoints(extractPointList(line), t)
		style := extractStyle(line)
		return fmt.Sprintf(`<polygon points="%s" %s/>`, points, style)

	case strings.HasPrefix(line, "PL["):
		points := transformPoints(extractPointList(line), t)
		style := extractStyle(line)
		return fmt.Sprintf(`<polyline points="%s" %s/>`, points, style)

	default:
		return fmt.Sprintf("<!-- Unknown line: %s -->", line)
	}
}
