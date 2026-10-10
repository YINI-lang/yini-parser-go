// Package yini parses YINI configuration documents.
//
// Parsing is lenient by default. Set Options.Mode to ModeStrict to enforce
// strict document rules. Parse results include diagnostics even when an error
// is returned.
package yini

import (
	"fmt"
)

// Mode selects the YINI validation rules.
type Mode string

const (
	ModeLenient Mode = "lenient"
	ModeStrict  Mode = "strict"
)

// Options controls parsing. The zero value selects lenient mode.
type Options struct {
	Mode Mode
}

// Position identifies a location in the original input. Lines and columns
// start at one.
type Position struct {
	Line   int
	Column int
}

// Severity is the level of a parse diagnostic.
type Severity string

const (
	SeverityWarning Severity = "warning"
	SeverityError   Severity = "error"
)

// Diagnostic describes a problem found while parsing or validating input.
type Diagnostic struct {
	Severity Severity
	Code     string
	Message  string
	Position Position
}

// Member is a key and its parsed value. Values are strings, int64, float64,
// bool, nil, []any, or map[string]any.
type Member struct {
	Name     string
	Value    any
	Position Position
}

// Section is an explicit YINI section and its nested content.
type Section struct {
	Name     string
	Level    int
	Position Position
	Members  []*Member
	Sections []*Section
}

// Document is the structured result of parsing a YINI document. Members at
// this level are lenient-mode orphan members.
type Document struct {
	Members  []*Member
	Sections []*Section
}

// Values returns a deep copy of the document's application-facing values.
// Orphan members and top-level sections appear beside one another.
func (d *Document) Values() map[string]any {
	values := make(map[string]any)
	if d == nil {
		return values
	}
	for _, member := range d.Members {
		values[member.Name] = copyValue(member.Value)
	}
	for _, section := range d.Sections {
		values[section.Name] = sectionValues(section)
	}
	return values
}

func sectionValues(section *Section) map[string]any {
	values := make(map[string]any)
	for _, member := range section.Members {
		values[member.Name] = copyValue(member.Value)
	}
	for _, child := range section.Sections {
		values[child.Name] = sectionValues(child)
	}
	return values
}

func copyValue(value any) any {
	switch value := value.(type) {
	case []any:
		copy := make([]any, len(value))
		for i, item := range value {
			copy[i] = copyValue(item)
		}
		return copy
	case map[string]any:
		copy := make(map[string]any, len(value))
		for key, item := range value {
			copy[key] = copyValue(item)
		}
		return copy
	default:
		return value
	}
}

// Result contains parsed data and all diagnostics collected from the input.
// Document may be partial when errors were reported.
type Result struct {
	Document    *Document
	Diagnostics []Diagnostic
}

// HasErrors reports whether the result contains an error diagnostic.
func (r Result) HasErrors() bool {
	for _, diagnostic := range r.Diagnostics {
		if diagnostic.Severity == SeverityError {
			return true
		}
	}
	return false
}

// ParseError reports a syntax or validation failure. The corresponding
// Result retains the complete list of diagnostics.
type ParseError struct {
	Diagnostic Diagnostic
}

func (e *ParseError) Error() string {
	if e == nil {
		return "YINI parse error"
	}
	position := e.Diagnostic.Position
	if position.Line > 0 {
		return fmt.Sprintf("YINI line %d, column %d: %s", position.Line, position.Column, e.Diagnostic.Message)
	}
	return "YINI: " + e.Diagnostic.Message
}
