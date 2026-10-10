package yini

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type stringDecodeError struct {
	offset  int // Byte offset within the token text.
	message string
}

func (e *stringDecodeError) Error() string { return e.message }

func decodeString(text string) (string, error) {
	if text == "" {
		return "", &stringDecodeError{0, "Empty string token"}
	}
	prefixLength := 0
	classic := false
	if len(text) > 1 && (text[0] == 'C' || text[0] == 'c' || text[0] == 'R' || text[0] == 'r') {
		prefixLength = 1
		classic = text[0] == 'C' || text[0] == 'c'
	}
	literal := text[prefixLength:]
	delimiterLength := 1
	triple := strings.HasPrefix(literal, `"""`)
	if triple {
		delimiterLength = 3
	}
	if len(literal) < delimiterLength*2 {
		return "", &stringDecodeError{0, "String is missing its closing quote"}
	}
	if triple {
		if !strings.HasSuffix(literal, `"""`) {
			return "", &stringDecodeError{0, "String is missing its closing triple quote"}
		}
	} else if (literal[0] != '\'' && literal[0] != '"') || literal[len(literal)-1] != literal[0] {
		return "", &stringDecodeError{0, "String is missing its closing quote"}
	}
	start := prefixLength + delimiterLength
	inner := text[start : len(text)-delimiterLength]
	var decoded strings.Builder
	for index := 0; index < len(inner); {
		char, width := utf8.DecodeRuneInString(inner[index:])
		if char == utf8.RuneError && width == 1 {
			return "", &stringDecodeError{start + index, "String contains invalid UTF-8"}
		}
		if char < 0x20 && char != '\t' && !(triple && (char == '\r' || char == '\n')) {
			return "", &stringDecodeError{start + index, "Literal control characters are not allowed in strings"}
		}
		if !classic || char != '\\' {
			decoded.WriteRune(char)
			index += width
			continue
		}
		if index+1 >= len(inner) {
			return "", &stringDecodeError{start + index, "String ends with an incomplete escape"}
		}
		escape := inner[index+1]
		if replacement, ok := simpleEscapes[escape]; ok {
			decoded.WriteRune(replacement)
			index += 2
			continue
		}
		base, count := 16, 0
		switch escape {
		case 'x':
			count = 2
		case 'u':
			count = 4
		case 'U':
			count = 8
		case 'o':
			base = 8
			for next := index + 2; next < len(inner) && count < 3 && inner[next] >= '0' && inner[next] <= '7'; next++ {
				count++
			}
			if count == 0 || (index+2+count < len(inner) && inner[index+2+count] >= '0' && inner[index+2+count] <= '9') {
				return "", &stringDecodeError{start + index, "Invalid octal escape; use one to three digits from 0 to 7"}
			}
		default:
			return "", &stringDecodeError{start + index, fmt.Sprintf("Unsupported escape \\%c", escape)}
		}
		if index+2+count > len(inner) {
			return "", &stringDecodeError{start + index, fmt.Sprintf("Escape \\%c needs %d digits", escape, count)}
		}
		digits := inner[index+2 : index+2+count]
		value, err := strconv.ParseUint(digits, base, 32)
		if err != nil || (escape == 'o' && value > 255) || value > utf8.MaxRune || (value >= 0xD800 && value <= 0xDFFF) {
			return "", &stringDecodeError{start + index, fmt.Sprintf("Invalid escape \\%c%s", escape, digits)}
		}
		decoded.WriteRune(rune(value))
		index += 2 + count
	}
	return decoded.String(), nil
}

var simpleEscapes = map[byte]rune{
	'\\': '\\', '\'': '\'', '"': '"', '/': '/', '0': 0, '?': '?',
	'a': '\a', 'b': '\b', 'f': '\f', 'n': '\n', 'r': '\r', 't': '\t', 'v': '\v',
}
