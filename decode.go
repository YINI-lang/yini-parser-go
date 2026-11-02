package yini

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// DecodeError reports a failure to decode a YINI value into a Go type.
type DecodeError struct {
	Path     string
	Position Position
	Message  string
}

func (e *DecodeError) Error() string {
	if e == nil {
		return "YINI decode error"
	}
	message := "YINI decode error"
	if e.Position.Line > 0 {
		message = fmt.Sprintf("YINI line %d, column %d", e.Position.Line, e.Position.Column)
	}
	if e.Path != "" {
		message += " at " + e.Path
	}
	return message + ": " + e.Message
}

// Decode copies the document root into dst, which must be a non-nil pointer
// to a struct. Orphan members and top-level sections are peers.
func (d *Document) Decode(dst any) error {
	if d == nil {
		return &DecodeError{Message: "cannot decode a nil document"}
	}
	return decodeRoot(dst, d.Members, d.Sections, "")
}

// Decode copies the section's members and child sections into dst, which
// must be a non-nil pointer to a struct.
func (s *Section) Decode(dst any) error {
	if s == nil {
		return &DecodeError{Message: "cannot decode a nil section"}
	}
	return decodeRoot(dst, s.Members, s.Sections, decodePath("", s.Name))
}

func decodeRoot(dst any, members []*Member, sections []*Section, path string) error {
	value := reflect.ValueOf(dst)
	if !value.IsValid() || value.Kind() != reflect.Pointer || value.IsNil() || value.Elem().Kind() != reflect.Struct {
		return &DecodeError{Message: "destination must be a non-nil pointer to a struct"}
	}
	if err := validateDecodeType(value.Elem().Type(), make(map[reflect.Type]bool)); err != nil {
		return err
	}
	fresh := reflect.New(value.Elem().Type()).Elem()
	if err := decodeEntries(sectionEntries(members, sections), fresh, path); err != nil {
		return err
	}
	value.Elem().Set(fresh)
	return nil
}

func validateDecodeType(typ reflect.Type, seen map[reflect.Type]bool) error {
	if seen[typ] {
		return nil
	}
	seen[typ] = true
	switch typ.Kind() {
	case reflect.Pointer, reflect.Slice:
		return validateDecodeType(typ.Elem(), seen)
	case reflect.Map:
		if typ.Key() != reflect.TypeOf("") {
			return &DecodeError{Message: fmt.Sprintf("Go type %s needs string map keys", typ)}
		}
		return validateDecodeType(typ.Elem(), seen)
	case reflect.Interface:
		if typ.NumMethod() != 0 {
			return &DecodeError{Message: fmt.Sprintf("Go type %s is a non-empty interface", typ)}
		}
	case reflect.Struct:
		names := make(map[string]string)
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			name, tagged := field.Tag.Lookup("yini")
			if tagged && name == "-" {
				continue
			}
			if field.PkgPath != "" {
				continue
			}
			if field.Anonymous {
				return &DecodeError{Message: fmt.Sprintf("Go field %s.%s is embedded; name it explicitly or use yini:\"-\"", typ, field.Name)}
			}
			if !tagged {
				name = field.Name
			} else if strings.Contains(name, ",") {
				return &DecodeError{Message: fmt.Sprintf("Go field %s.%s has unsupported yini tag options", typ, field.Name)}
			}
			if previous, exists := names[name]; exists {
				return &DecodeError{Message: fmt.Sprintf("Go fields %s.%s and %s.%s both map to YINI name %q", typ, previous, typ, field.Name, name)}
			}
			names[name] = field.Name
			if err := validateDecodeType(field.Type, seen); err != nil {
				return err
			}
		}
	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float64:
	default:
		return &DecodeError{Message: fmt.Sprintf("Go type %s is not supported for YINI decoding", typ)}
	}
	return nil
}

type decodeEntry struct {
	name     string
	value    any
	section  *Section
	position Position
}

func sectionEntries(members []*Member, sections []*Section) []decodeEntry {
	entries := make([]decodeEntry, 0, len(members)+len(sections))
	for _, member := range members {
		if member != nil {
			entries = append(entries, decodeEntry{name: member.Name, value: member.Value, position: member.Position})
		}
	}
	for _, section := range sections {
		if section != nil {
			entries = append(entries, decodeEntry{name: section.Name, section: section, position: section.Position})
		}
	}
	sort.SliceStable(entries, func(i, j int) bool {
		left, right := entries[i].position, entries[j].position
		if left.Line != right.Line {
			return left.Line < right.Line
		}
		return left.Column < right.Column
	})
	return entries
}

func decodeEntries(entries []decodeEntry, dst reflect.Value, path string) error {
	switch dst.Kind() {
	case reflect.Struct:
		fields := make(map[string]int)
		for i := 0; i < dst.NumField(); i++ {
			field := dst.Type().Field(i)
			name, tagged := field.Tag.Lookup("yini")
			if tagged && name == "-" || field.PkgPath != "" {
				continue
			}
			if !tagged {
				name = field.Name
			}
			fields[name] = i
		}
		for _, entry := range entries {
			entryPath := decodePath(path, entry.name)
			index, exists := fields[entry.name]
			if !exists {
				return &DecodeError{Path: entryPath, Position: entry.position, Message: fmt.Sprintf("YINI %s %q has no matching field in Go %s", entryType(entry), entry.name, dst.Type())}
			}
			if err := decodeEntryValue(entry, dst.Field(index), entryPath); err != nil {
				return err
			}
		}
	case reflect.Map:
		result := reflect.MakeMapWithSize(dst.Type(), len(entries))
		for _, entry := range entries {
			entryPath := decodePath(path, entry.name)
			value := reflect.New(dst.Type().Elem()).Elem()
			if err := decodeEntryValue(entry, value, entryPath); err != nil {
				return err
			}
			result.SetMapIndex(reflect.ValueOf(entry.name), value)
		}
		dst.Set(result)
	default:
		return &DecodeError{Path: path, Message: fmt.Sprintf("cannot decode YINI object into Go %s", dst.Type())}
	}
	return nil
}

func decodeEntryValue(entry decodeEntry, dst reflect.Value, path string) error {
	if entry.section != nil {
		return decodeSection(entry.section, dst, path)
	}
	return decodeValue(entry.value, dst, path, entry.position)
}

func decodeSection(section *Section, dst reflect.Value, path string) error {
	if dst.Kind() == reflect.Pointer {
		dst.Set(reflect.New(dst.Type().Elem()))
		return decodeSection(section, dst.Elem(), path)
	}
	if dst.Kind() == reflect.Interface {
		dst.Set(reflect.ValueOf(sectionValues(section)))
		return nil
	}
	if dst.Kind() != reflect.Struct && dst.Kind() != reflect.Map {
		return &DecodeError{Path: path, Position: section.Position, Message: fmt.Sprintf("cannot decode YINI section into Go %s; need a struct or map[string]T", dst.Type())}
	}
	return decodeEntries(sectionEntries(section.Members, section.Sections), dst, path)
}

func decodeValue(value any, dst reflect.Value, path string, position Position) error {
	if value == nil {
		switch dst.Kind() {
		case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Interface:
			dst.SetZero()
			return nil
		default:
			return &DecodeError{Path: path, Position: position, Message: fmt.Sprintf("cannot decode YINI null into Go %s", dst.Type())}
		}
	}
	if dst.Kind() == reflect.Pointer {
		dst.Set(reflect.New(dst.Type().Elem()))
		return decodeValue(value, dst.Elem(), path, position)
	}
	if dst.Kind() == reflect.Interface {
		switch value.(type) {
		case string, bool, int64, float64, []any, map[string]any:
			dst.Set(reflect.ValueOf(copyValue(value)))
			return nil
		}
	}
	switch source := value.(type) {
	case string:
		if dst.Kind() == reflect.String {
			dst.SetString(source)
			return nil
		}
	case bool:
		if dst.Kind() == reflect.Bool {
			dst.SetBool(source)
			return nil
		}
	case int64:
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !dst.OverflowInt(source) {
				dst.SetInt(source)
				return nil
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if source >= 0 && !dst.OverflowUint(uint64(source)) {
				dst.SetUint(uint64(source))
				return nil
			}
		}
	case float64:
		if dst.Kind() == reflect.Float64 {
			dst.SetFloat(source)
			return nil
		}
	case []any:
		if dst.Kind() == reflect.Slice {
			result := reflect.MakeSlice(dst.Type(), len(source), len(source))
			for i, item := range source {
				itemPath := fmt.Sprintf("%s[%d]", path, i)
				if err := decodeValue(item, result.Index(i), itemPath, position); err != nil {
					return err
				}
			}
			dst.Set(result)
			return nil
		}
	case map[string]any:
		if dst.Kind() == reflect.Struct || dst.Kind() == reflect.Map {
			keys := make([]string, 0, len(source))
			for key := range source {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			entries := make([]decodeEntry, 0, len(keys))
			for _, key := range keys {
				entries = append(entries, decodeEntry{name: key, value: source[key], position: position})
			}
			return decodeEntries(entries, dst, path)
		}
	}
	return &DecodeError{Path: path, Position: position, Message: fmt.Sprintf("cannot decode YINI %s into Go %s", valueType(value), dst.Type())}
}

func entryType(entry decodeEntry) string {
	if entry.section != nil {
		return "section"
	}
	return valueType(entry.value)
}

func valueType(value any) string {
	switch value.(type) {
	case nil:
		return "null"
	case string:
		return "string"
	case bool:
		return "boolean"
	case int64:
		return "integer"
	case float64:
		return "float"
	case []any:
		return "list"
	case map[string]any:
		return "inline object"
	default:
		return fmt.Sprintf("value of type %T", value)
	}
}

func decodePath(parent, name string) string {
	plain := name != ""
	for i, char := range name {
		if !(char == '_' || char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || i > 0 && char >= '0' && char <= '9') {
			plain = false
			break
		}
	}
	if plain {
		if parent == "" {
			return name
		}
		return parent + "." + name
	}
	return parent + "[" + strconv.Quote(name) + "]"
}
