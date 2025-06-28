package converter

import (
	"fmt"
	"strings"

	"github.com/prabinpanta0/VectorFormatBridge/pkg/egf"
	"github.com/prabinpanta0/VectorFormatBridge/pkg/svg"
	"github.com/prabinpanta0/VectorFormatBridge/pkg/transform"
)

// SVGToEGF converts an SVG file to EGF format
func SVGToEGF(svgFile string, egfFile string) error {
	svgData, err := svg.ParseSVG(svgFile)
	if err != nil {
		return fmt.Errorf("failed to parse SVG: %w", err)
	}

	egfContent := fmt.Sprintf("M(%s,%s,#fff)\n", svgData.Width, svgData.Height)

	entityMap := map[string]string{}
	entityCount := 1

	// Helper to handle entities
	addEntity := func(def string) string {
		for k, v := range entityMap {
			if v == def {
				return k // reuse
			}
		}
		id := fmt.Sprintf("#%02d", entityCount)
		entityCount++
		entityMap[id] = def
		return id
	}

	// Process basic elements
	for _, r := range svgData.Rects {
		cmd := fmt.Sprintf("R(%s,%s,%s,%s) S(%s,%s)", r.X, r.Y, r.Width, r.Height, colorOrDefault(r.Stroke, "#000"), colorOrDefault(r.Fill, "#none"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	for _, c := range svgData.Circles {
		cmd := fmt.Sprintf("C(%s,%s,%s) S(%s,%s)", c.Cx, c.Cy, c.R, colorOrDefault(c.Stroke, "#000"), colorOrDefault(c.Fill, "#none"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	for _, l := range svgData.Lines {
		cmd := fmt.Sprintf("L(%s,%s,%s,%s) S(%s)", l.X1, l.Y1, l.X2, l.Y2, colorOrDefault(l.Stroke, "#000"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	for _, p := range svgData.Paths {
		cmd := fmt.Sprintf("P[%s] S(%s,%s)", sanitizePath(p.D), colorOrDefault(p.Stroke, "#000"), colorOrDefault(p.Fill, "#none"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	for _, e := range svgData.Ellipses {
		cmd := fmt.Sprintf("E(%s,%s,%s,%s) S(%s,%s)", e.Cx, e.Cy, e.Rx, e.Ry, colorOrDefault(e.Stroke, "#000"), colorOrDefault(e.Fill, "#none"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	for _, poly := range svgData.Polygons {
		cmd := fmt.Sprintf("PG[%s] S(%s,%s)", sanitizePoints(poly.Points), colorOrDefault(poly.Stroke, "#000"), colorOrDefault(poly.Fill, "#none"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	for _, pl := range svgData.Polylines {
		cmd := fmt.Sprintf("PL[%s] S(%s)", sanitizePoints(pl.Points), colorOrDefault(pl.Stroke, "#000"))
		id := addEntity(cmd)
		egfContent += fmt.Sprintf("CALL%s T(0,0,1,0)\n", id)
	}

	// Write entities at the top
	entityDefs := ""
	for id, def := range entityMap {
		entityDefs += fmt.Sprintf("H%s = %s\n", id, def)
	}
	egfContent = entityDefs + egfContent

	return egf.WriteEGF(egfFile, egfContent)
}

// EGFToSVG converts an EGF file to SVG format
func EGFToSVG(egfFile string, svgFile string) error {
	egfContent, err := egf.ReadEGF(egfFile)
	if err != nil {
		return fmt.Errorf("failed to read EGF: %w", err)
	}

	lines := strings.Split(egfContent, "\n")
	svgContent := `<svg xmlns="http://www.w3.org/2000/svg" width="800" height="600">` + "\n"

	entityMap := make(map[string]string) // H# entities

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		switch {
		case strings.HasPrefix(line, "M("):
			size := extractParams(line)
			if len(size) >= 2 {
				svgContent = fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%s" height="%s">`, size[0], size[1]) + "\n"
			}

		case strings.HasPrefix(line, "H#"):
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				id := strings.TrimSpace(strings.TrimPrefix(parts[0], "H"))
				entityMap[id] = strings.TrimSpace(parts[1])
			}

		case strings.HasPrefix(line, "CALL#"):
			parts := strings.SplitN(line, " ", 2)
			id := strings.TrimPrefix(parts[0], "CALL")
			entity, exists := entityMap[id]
			if exists {
				t := transform.NewTransform()
				if len(parts) > 1 {
					t = transform.ParseTransform(parts[1])
				}
				content := renderEntity(entity, t)
				svgContent += content + "\n"
			}

		case strings.HasPrefix(line, "G["):
			content := extractGroupContent(line)
			svgContent += content + "\n"

		default:
			svgContent += renderLine(line, transform.NewTransform()) + "\n"
		}
	}

	svgContent += "</svg>"

	return svg.WriteSVG(svgFile, svgContent)
}

// EGFToEGFB converts EGF to binary EGFB format
func EGFToEGFB(egfFile string, egfbFile string) error {
	egfContent, err := egf.ReadEGF(egfFile)
	if err != nil {
		return fmt.Errorf("failed to read EGF: %w", err)
	}

	return egf.EncodeToEGFB(egfContent, egfbFile)
}

// EGFBToEGF converts binary EGFB to EGF format
func EGFBToEGF(egfbFile string, egfFile string) error {
	egfContent, err := egf.DecodeFromEGFB(egfbFile)
	if err != nil {
		return fmt.Errorf("failed to decode EGFB: %w", err)
	}

	return egf.WriteEGF(egfFile, egfContent)
}

// Helper functions

// colorOrDefault returns the color if it's not empty, otherwise returns the default
func colorOrDefault(color, defaultColor string) string {
	if color == "" {
		return defaultColor
	}
	return color
}

// sanitizePath cleans up SVG path data for EGF format
func sanitizePath(path string) string {
	return strings.TrimSpace(strings.ReplaceAll(path, "  ", " "))
}

// sanitizePoints cleans up SVG points data for EGF format
func sanitizePoints(points string) string {
	return strings.TrimSpace(strings.ReplaceAll(points, "  ", " "))
}
