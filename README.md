# VectorFormatBridge

A powerful command-line tool for converting between SVG (Scalable Vector Graphics) and EGF (Enhanced Graphics Format), with support for binary compression.

## 🎯 What is EGF?

**Enhanced Graphics Format (EGF)** is a custom text-based vector graphics format designed for efficiency and simplicity. Unlike SVG's XML structure, EGF uses a more compact command-based syntax that's both human-readable and machine-optimized.

### EGF Format Features

- **Text-based**: Human-readable commands for easy debugging and manual editing
- **Compact syntax**: More concise than SVG while maintaining full vector capabilities  
- **Entity reuse**: Supports defining reusable graphics entities to reduce file size
- **Transform support**: Built-in transformation matrices for scaling, rotation, and translation
- **Binary compression**: EGFB format provides compressed binary storage
- **Shape primitives**: Rectangles, circles, lines, paths, ellipses, polygons, and polylines

### EGF Syntax Overview

```
M(width,height,background)           # Canvas/viewport definition
R(x,y,width,height) S(stroke,fill)   # Rectangle with styling  
C(cx,cy,radius) S(stroke,fill)       # Circle with styling
L(x1,y1,x2,y2) S(stroke)            # Line with stroke
P[path_data] S(stroke,fill)          # Path element
E(cx,cy,rx,ry) S(stroke,fill)        # Ellipse with styling
PG[points] S(stroke,fill)            # Polygon with styling
PL[points] S(stroke)                 # Polyline with stroke
H#01 = R(10,10,50,50) S(#000,#f00)  # Entity definition
CALL#01 T(100,100,1.5,45)           # Entity call with transform
```

## 🚀 Supported Conversions

| From | To | Description |
|------|----|-----------| 
| SVG  | EGF  | Convert SVG files to Enhanced Graphics Format |
| EGF  | SVG  | Convert EGF files back to standard SVG |
| EGF  | EGFB | Encode EGF to compressed binary format |
| EGFB | EGF  | Decode binary EGFB back to text EGF |

## 📦 Installation

### Prerequisites
- Go 1.19 or later

### Build from Source
```bash
git clone https://github.com/prabinpanta0/VectorFormatBridge.git
cd VectorFormatBridge
go build -o vectorformatbridge cmd/vectorformatbridge/main.go
```

### Install with Go
```bash
go install github.com/yourusername/VectorFormatBridge/cmd/vectorformatbridge@latest
```

## 🛠️ Usage

### Basic Commands

```bash
# Convert SVG to EGF
vectorformatbridge svg2egf input.svg output.egf

# Convert EGF to SVG  
vectorformatbridge egf2svg input.egf output.svg

# Compress EGF to binary format
vectorformatbridge egf2egfb input.egf output.egfb

# Decompress binary format back to EGF
vectorformatbridge egfb2egf input.egfb output.egf

# Run demo with sample files
vectorformatbridge demo
```

### Usage Examples

#### Converting an SVG file
```bash
vectorformatbridge svg2egf logo.svg logo.egf
```

#### Converting back to SVG
```bash  
vectorformatbridge egf2svg logo.egf logo_converted.svg
```

#### Creating compressed binary
```bash
vectorformatbridge egf2egfb graphics.egf graphics.egfb
```

## 📊 Format Comparison

| Feature | SVG | EGF | EGFB |
|---------|-----|-----|------|
| File Type | XML Text | Custom Text | Binary |
| Human Readable | ✅ | ✅ | ❌ |
| File Size | Large | Medium | Small |
| Parse Speed | Slow | Fast | Fastest |
| Browser Support | ✅ | ❌ | ❌ |
| Entity Reuse | Limited | ✅ | ✅ |
| Transform Support | ✅ | ✅ | ✅ |

## 🎨 Supported SVG Elements

- **Basic Shapes**: `<rect>`, `<circle>`, `<line>`, `<ellipse>`
- **Complex Shapes**: `<path>`, `<polygon>`, `<polyline>`
- **Styling**: `fill`, `stroke` attributes
- **Transforms**: Translation, scaling, rotation (via EGF transform syntax)

## 📁 Project Structure

```
VectorFormatBridge/
├── cmd/vectorformatbridge/     # Main application entry point
│   └── main.go
├── pkg/
│   ├── svg/                    # SVG parsing and generation
│   ├── egf/                    # EGF format handling  
│   ├── converter/              # Format conversion logic
│   └── transform/              # Transformation utilities
├── examples/                   # Example files and demos
├── README.md
└── go.mod
```

## 🔧 Advanced Features

### Entity System
EGF supports defining reusable graphics entities:
```
H#01 = R(0,0,10,10) S(#000,#f00)    # Define red square entity
CALL#01 T(50,50,2.0,0)              # Use entity at position with 2x scale
CALL#01 T(100,100,1.0,45)           # Use entity rotated 45 degrees  
```

### Transform Matrices
Apply transformations using `T(x,y,scale,rotate)` syntax:
- `x,y`: Translation coordinates
- `scale`: Uniform scaling factor
- `rotate`: Rotation angle in degrees

### Binary Compression
EGFB format provides significant file size reduction:
- Efficient binary encoding of commands
- Compressed coordinate data
- Optimized for parsing speed

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup
```bash
git clone https://github.com/prabinpanta/VectorFormatBridge.git
cd VectorFormatBridge
go mod tidy
go test ./...
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- SVG specification by W3C
- Go XML parsing capabilities
- Binary encoding techniques for graphics optimization

## 📚 EGF Format Specification

### Command Reference

| Command | Syntax | Description |
|---------|--------|-------------|
| M | `M(w,h,bg)` | Define canvas size and background |
| R | `R(x,y,w,h) S(stroke,fill)` | Rectangle |
| C | `C(cx,cy,r) S(stroke,fill)` | Circle |
| L | `L(x1,y1,x2,y2) S(stroke)` | Line |
| E | `E(cx,cy,rx,ry) S(stroke,fill)` | Ellipse |
| P | `P[data] S(stroke,fill)` | Path |
| PG | `PG[points] S(stroke,fill)` | Polygon |
| PL | `PL[points] S(stroke)` | Polyline |
| H | `H#id = command` | Entity definition |
| CALL | `CALL#id T(x,y,s,r)` | Entity instantiation |

### Color Format
- Hex colors: `#RGB` or `#RRGGBB`
- Named colors: `red`, `blue`, `green`, etc.
- Transparent: `#none`

### Coordinate System
- Origin (0,0) at top-left
- X increases rightward
- Y increases downward
- All coordinates in pixels (can be fractional)

---

*VectorFormatBridge - Bridging the gap between vector graphics formats*
```
: ¨·.·¨ :
 ` ·. 🦋
                  ╱|、                   
                (˚ˎ 。7  
                |、˜〵          
                じしˍ,)ノ           
```
