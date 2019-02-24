package pe

import (
	"bytes"
	"time"
)

// ### [ Helper functions ] ####################################################

// parseDateFromEpoch parses the given date in number of seconds since Epoch
// into a corresponding Go date.
func parseDateFromEpoch(date uint32) time.Time {
	return time.Unix(int64(date), 0)
}

// parseCString parses the given a NULL-terminated string into a corresponding
// Go string.
func parseCString(b []byte) string {
	pos := bytes.IndexByte(b, '\x00')
	if pos != -1 {
		b = b[:pos]
	}
	return string(b)
}

// parseCString parses a NULL-terminated string at the given address into a
// corresponding Go string.
func (file *File) parseCString(addr uint64) string {
	var buf []byte
	for {
		bs := file.ReadData(addr, 1)
		b := bs[0]
		if b == '\x00' {
			// Break at NULL-byte
			break
		}
		buf = append(buf, b)
		addr++
	}
	return string(buf)
}
