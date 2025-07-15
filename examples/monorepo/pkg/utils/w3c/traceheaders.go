//
//  traceheaders.go
//  newest_template
//
//  Created by karim-w on 12/07/2025.
//

package w3c

import (
	"fmt"
	"strings"
)

var (
	ErrInvalidTraceHeaderFormat = fmt.Errorf("invalid trace header format, expected '00-<trace_id>-<parent_id>-<flags>'")
	ErrInvalidTraceID           = fmt.Errorf("invalid trace_id, expected 32-character hexadecimal string")
	ErrInvalidParentID          = fmt.Errorf("invalid parent_id, expected 16-character hexadecimal string")
)

// ParseTraceHeaders parses a W3C trace header string and returns the
// trace ID, parent ID(old span_id), and a boolean indicating validity.
// The trace header is expected to be in the format "00-<trace_id>-<parent_id>-<flags>",
// where <trace_id> is a 32-character hexadecimal string and <parent_id> is a 16-character hexadecimal string.
// If the trace header is invalid, it returns false for is_valid and an error.
func ParseTraceHeaders(
	traceHeader string,
) (is_valid bool, trace_id string, parent_id string, err error) {
	parts := strings.Split(traceHeader, "-")
	if len(parts) != 4 {
		return false, "", "", ErrInvalidTraceHeaderFormat
	}

	trace_id = parts[1]
	parent_id = parts[2]

	if len(trace_id) != 32 || !isHex(trace_id) {
		return false, "", "", fmt.Errorf("invalid trace_id: %s", trace_id)
	}
	if len(parent_id) != 16 || !isHex(parent_id) {
		return false, "", "", fmt.Errorf("invalid parent_id: %s", parent_id)
	}

	return true, trace_id, parent_id, nil
}

func isHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
