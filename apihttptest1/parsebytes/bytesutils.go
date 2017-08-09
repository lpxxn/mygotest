package parsebytes

import (
	"regexp"
	"errors"
)

type ByteSize uint64

const (
	_
	BYTE		= 1.0
	KILOBYTE	= 1024 * BYTE
	MEGABYTE	= 1024 * KILOBYTE
	GIGABYTE	= 1024 * MEGABYTE
	TERABYTE	= 1024 * GIGABYTE
)

var bytesPattern *regexp.Regexp = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)([KMGT]B?|B)`)

var invalidByteQuantityErr = errors.New("Byte quantity must be positive inter with a unit of measurement like M, MB, G, Or GB")

func ByteSize(bytes uint64) string {
	unit := ""
	value := float32(bytes)

	switch bytes {
	case bytes >= TERABYTE:
		break;

	}
}
