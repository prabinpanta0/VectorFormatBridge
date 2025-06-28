package egf

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"strings"
)

// ErrInvalidEGFB is returned when the EGFB file format is invalid
var ErrInvalidEGFB = errors.New("invalid EGFB file format")

// WriteEGF writes EGF content to a file
func WriteEGF(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// ReadEGF reads EGF content from a file
func ReadEGF(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// EncodeToEGFB encodes EGF text to binary EGFB format
func EncodeToEGFB(egfContent string, egfbFile string) error {
	lines := strings.Split(egfContent, "\n")
	buf := new(bytes.Buffer)

	// Write header
	buf.WriteString("EGFB")

	// Process each line
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		code := getOpcode(line)
		buf.WriteByte(code)

		// Encode the entire line as a string for simplicity
		lineBytes := []byte(line)
		// Write length as 2 bytes (little endian) to handle longer strings
		binary.Write(buf, binary.LittleEndian, uint16(len(lineBytes)))
		buf.Write(lineBytes)
	}

	// End mark
	buf.WriteByte(0xFF)

	return ioutil.WriteFile(egfbFile, buf.Bytes(), 0644)
}

// DecodeFromEGFB decodes binary EGFB format back to EGF text
func DecodeFromEGFB(egfbFile string) (string, error) {
	data, err := ioutil.ReadFile(egfbFile)
	if err != nil {
		return "", err
	}

	pos := 0
	if len(data) < 4 || string(data[0:4]) != "EGFB" {
		return "", ErrInvalidEGFB
	}
	pos += 4

	egf := ""

	// Decode binary back to EGF by reading each encoded line as a string
	for pos < len(data) {
		op := data[pos]
		pos++
		if op == 0xFF {
			break
		}
		// Next two bytes: length of the line
		if pos+2 > len(data) {
			break
		}
		length := binary.LittleEndian.Uint16(data[pos : pos+2])
		pos += 2
		if pos+int(length) > len(data) {
			break
		}
		line := string(data[pos : pos+int(length)])
		pos += int(length)
		egf += line + "\n"
	}

	return egf, nil
}

// getOpcode returns the opcode for a given EGF command line
func getOpcode(line string) byte {
	switch {
	case strings.HasPrefix(line, "M("):
		return 0x01
	case strings.HasPrefix(line, "R("):
		return 0x02
	case strings.HasPrefix(line, "C("):
		return 0x03
	case strings.HasPrefix(line, "L("):
		return 0x04
	case strings.HasPrefix(line, "P["):
		return 0x05
	case strings.HasPrefix(line, "E("):
		return 0x06
	case strings.HasPrefix(line, "PG["):
		return 0x07
	case strings.HasPrefix(line, "PL["):
		return 0x08
	case strings.HasPrefix(line, "H#"):
		return 0x10
	case strings.HasPrefix(line, "CALL#"):
		return 0x11
	default:
		return 0x00
	}
}
