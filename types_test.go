package yini

import (
	"reflect"
	"testing"
)

func TestCopyValue(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name  string
		value any
	}{
		{"null", nil},
		{"exact integer", int64(9007199254740993)},
		{"float", 0.125},
		{"string", "cedar"},
		{"Boolean", true},
		{"empty list", []any{}},
		{"empty object", map[string]any{}},
		{"nested null and empty values", map[string]any{"items": []any{nil, []any{}, map[string]any{}}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			value := test.value

			// 2. Act.
			got := copyValue(value)

			// 3. Assert.
			if !reflect.DeepEqual(got, value) {
				t.Errorf("copy = %#v (%T), want %#v (%T)", got, got, value, value)
			}
		})
	}
}

func TestCopyValueIsolation(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name         string
		mutateSource bool
	}{
		{name: "mutate copy"},
		{name: "mutate source", mutateSource: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			source := map[string]any{
				"items": []any{map[string]any{"labels": []any{"cedar"}}, nil},
				"count": int64(9007199254740993),
			}
			want := map[string]any{
				"items": []any{map[string]any{"labels": []any{"cedar"}}, nil},
				"count": int64(9007199254740993),
			}

			// 2. Act.
			copied := copyValue(source).(map[string]any)
			changed, unchanged := copied, source
			if test.mutateSource {
				changed, unchanged = source, copied
			}
			changed["count"] = int64(1)
			items := changed["items"].([]any)
			items[1] = "changed"
			item := items[0].(map[string]any)
			item["added"] = true
			item["labels"].([]any)[0] = "changed"

			// 3. Assert.
			if !reflect.DeepEqual(unchanged, want) {
				t.Errorf("mutating one value changed the other: got %#v, want %#v", unchanged, want)
			}
		})
	}
}

func TestSectionValuesIsolation(t *testing.T) {
	// 1. Assign.
	section := &Section{
		Name:    "Depot",
		Members: []*Member{{Name: "labels", Value: []any{"cedar"}}},
		Sections: []*Section{{
			Name: "Bay",
			Members: []*Member{{Name: "details", Value: map[string]any{
				"lanes": []any{int64(2)},
			}}},
		}},
	}
	want := map[string]any{
		"labels": []any{"cedar"},
		"Bay":    map[string]any{"details": map[string]any{"lanes": []any{int64(2)}}},
	}
	values := sectionValues(section)
	if !reflect.DeepEqual(values, want) {
		t.Fatalf("section values = %#v, want %#v", values, want)
	}

	// 2. Act.
	values["labels"].([]any)[0] = "changed"
	values["Bay"].(map[string]any)["details"].(map[string]any)["lanes"].([]any)[0] = int64(99)
	fresh := sectionValues(section)

	// 3. Assert.
	if !reflect.DeepEqual(fresh, want) {
		t.Errorf("section storage changed: got %#v, want %#v", fresh, want)
	}
}
