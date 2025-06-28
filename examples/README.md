# Examples

This directory contains example files demonstrating the VectorFormatBridge converter.

## Sample Files

- `sample.svg` - A sample SVG file with various shapes
- `sample.egf` - The same graphics in EGF format
- `README.md` - This file

## Running Examples

```bash
# Convert the sample SVG to EGF
vectorformatbridge svg2egf examples/sample.svg output.egf

# Convert EGF back to SVG
vectorformatbridge egf2svg examples/sample.egf output.svg

# Create binary version
vectorformatbridge egf2egfb examples/sample.egf output.egfb
```
