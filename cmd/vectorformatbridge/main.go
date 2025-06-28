package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/prabinpanta0/VectorFormatBridge/pkg/converter"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "svg2egf":
		if len(os.Args) != 4 {
			fmt.Println("Usage: vectorformatbridge svg2egf <input.svg> <output.egf>")
			return
		}
		err := converter.SVGToEGF(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("Error converting SVG to EGF: %v\n", err)
			return
		}
		fmt.Println("Converted SVG to EGF successfully.")

	case "egf2svg":
		if len(os.Args) != 4 {
			fmt.Println("Usage: vectorformatbridge egf2svg <input.egf> <output.svg>")
			return
		}
		err := converter.EGFToSVG(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("Error converting EGF to SVG: %v\n", err)
			return
		}
		fmt.Println("Converted EGF to SVG successfully.")

	case "egf2egfb":
		if len(os.Args) != 4 {
			fmt.Println("Usage: vectorformatbridge egf2egfb <input.egf> <output.egfb>")
			return
		}
		err := converter.EGFToEGFB(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("Error encoding EGF to EGFB: %v\n", err)
			return
		}
		fmt.Println("Encoded EGF to EGFB successfully.")

	case "egfb2egf":
		if len(os.Args) != 4 {
			fmt.Println("Usage: vectorformatbridge egfb2egf <input.egfb> <output.egf>")
			return
		}
		err := converter.EGFBToEGF(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("Error decoding EGFB to EGF: %v\n", err)
			return
		}
		fmt.Println("Decoded EGFB to EGF successfully.")

	case "demo":
		runDemo()

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("VectorFormatBridge - Bridge between vector graphics formats")
	fmt.Println()
	fmt.Println("Format explanations:")
	fmt.Println("  SVG  - Scalable Vector Graphics (XML-based)")
	fmt.Println("  EGF  - Enhanced Graphics Format (text-based intermediate format)")
	fmt.Println("  EGFB - EGF Binary (compressed binary version of EGF)")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  vectorformatbridge svg2egf <input.svg> <output.egf>   - Convert SVG to EGF")
	fmt.Println("  vectorformatbridge egf2svg <input.egf> <output.svg>   - Convert EGF to SVG")
	fmt.Println("  vectorformatbridge egf2egfb <input.egf> <output.egfb> - Encode EGF to binary EGFB")
	fmt.Println("  vectorformatbridge egfb2egf <input.egfb> <output.egf> - Decode EGFB back to EGF")
	fmt.Println("  vectorformatbridge demo                               - Run demo with sample files")
	fmt.Println()
	fmt.Println("Note: EGFB is a binary/compressed version of EGF for efficient storage.")
}

func runDemo() {
	fmt.Println("Running VectorFormatBridge demo...")

	// Create a sample SVG file
	sampleSVG := `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300">
  <rect x="50" y="50" width="100" height="75" fill="#ff0000" stroke="#000000"/>
  <circle cx="200" cy="150" r="40" fill="#00ff00" stroke="#000000"/>
  <line x1="300" y1="50" x2="350" y2="100" stroke="#0000ff"/>
  <ellipse cx="150" cy="200" rx="30" ry="20" fill="#ffff00" stroke="#000000"/>
  <polygon points="250,200 270,220 250,240 230,220" fill="#ff00ff" stroke="#000000"/>
</svg>`

	// Write sample SVG
	err := ioutil.WriteFile("demo.svg", []byte(sampleSVG), 0644)
	if err != nil {
		fmt.Println("Error creating demo.svg:", err)
		return
	}
	fmt.Println("Created demo.svg")

	// Convert SVG to EGF
	err = converter.SVGToEGF("demo.svg", "demo.egf")
	if err != nil {
		fmt.Printf("Error converting SVG to EGF: %v\n", err)
		return
	}
	fmt.Println("Converted demo.svg to demo.egf")

	// Convert EGF back to SVG
	err = converter.EGFToSVG("demo.egf", "demo_converted.svg")
	if err != nil {
		fmt.Printf("Error converting EGF to SVG: %v\n", err)
		return
	}
	fmt.Println("Converted demo.egf to demo_converted.svg")

	// Encode EGF to binary
	err = converter.EGFToEGFB("demo.egf", "demo.egfb")
	if err != nil {
		fmt.Printf("Error encoding EGF to EGFB: %v\n", err)
		return
	}
	fmt.Println("Encoded demo.egf to demo.egfb")

	// Decode binary back to EGF
	err = converter.EGFBToEGF("demo.egfb", "demo_decoded.egf")
	if err != nil {
		fmt.Printf("Error decoding EGFB to EGF: %v\n", err)
		return
	}
	fmt.Println("Decoded demo.egfb to demo_decoded.egf")

	fmt.Println()
	fmt.Println("Demo completed! Files created:")
	fmt.Println("- demo.svg (original)")
	fmt.Println("- demo.egf (EGF format)")
	fmt.Println("- demo_converted.svg (converted back from EGF)")
	fmt.Println("- demo.egfb (binary format)")
	fmt.Println("- demo_decoded.egf (decoded from binary)")
}
