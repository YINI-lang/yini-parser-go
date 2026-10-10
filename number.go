package yini

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// parseNumber converts every supported literal without routing integers through
// floating point. The second result is the canonical form used by lenient
// string concatenation.
func parseNumber(text string) (any, string, error) {
	clean := strings.ReplaceAll(text, "_", "")
	sign := ""
	if strings.HasPrefix(clean, "+") || strings.HasPrefix(clean, "-") {
		sign, clean = clean[:1], clean[1:]
	}
	lower := strings.ToLower(clean)
	base := 10
	digits := clean
	switch {
	case strings.HasPrefix(lower, "0b"):
		base, digits = 2, clean[2:]
	case strings.HasPrefix(lower, "%"):
		base, digits = 2, clean[1:]
	case strings.HasPrefix(lower, "0o"):
		base, digits = 8, clean[2:]
	case strings.HasPrefix(lower, "0z"):
		base, digits = 12, clean[2:]
		digits = strings.NewReplacer("x", "a", "X", "A", "e", "b", "E", "B").Replace(digits)
	case strings.HasPrefix(lower, "0x"):
		base, digits = 16, clean[2:]
	case strings.HasPrefix(lower, "hex:"):
		base, digits = 16, clean[4:]
	}
	if base != 10 || (!strings.ContainsAny(clean, ".eE")) {
		value, err := strconv.ParseInt(sign+digits, base, 64)
		if err != nil {
			return nil, "", fmt.Errorf("integer %q is outside the signed 64-bit range or is invalid", text)
		}
		return value, strconv.FormatInt(value, 10), nil
	}
	value, err := strconv.ParseFloat(sign+clean, 64)
	if err != nil && !errors.Is(err, strconv.ErrRange) {
		return nil, "", fmt.Errorf("invalid floating-point number %q", text)
	}
	if math.IsInf(value, 0) {
		return nil, "", fmt.Errorf("floating-point number %q is too large", text)
	}
	if value == 0 && containsNonzeroDigit(clean) {
		return nil, "", fmt.Errorf("floating-point number %q is too small to represent", text)
	}
	return value, strconv.FormatFloat(value, 'g', -1, 64), nil
}

func containsNonzeroDigit(text string) bool {
	for _, char := range text {
		if char == 'e' || char == 'E' {
			break
		}
		if char >= '1' && char <= '9' {
			return true
		}
	}
	return false
}
