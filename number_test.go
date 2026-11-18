package yini

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestParseNumber(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name      string
		literal   string
		want      any
		canonical string
	}{
		{"decimal separators", "9_007_199_254_740_993", int64(9007199254740993), "9007199254740993"},
		{"decimal minimum", "-9223372036854775808", int64(math.MinInt64), "-9223372036854775808"},
		{"decimal maximum", "9223372036854775807", int64(math.MaxInt64), "9223372036854775807"},
		{"positive sign", "+42", int64(42), "42"},
		{"integer negative zero", "-0", int64(0), "-0"},
		{"binary separators", "0b_1010_0110", int64(166), "166"},
		{"signed binary", "-0B101", int64(-5), "-5"},
		{"legacy binary", "%1101", int64(13), "13"},
		{"negative legacy binary", "-%1101", int64(-13), "-13"},
		{"binary maximum", "0b" + strings.Repeat("1", 63), int64(math.MaxInt64), "9223372036854775807"},
		{"binary minimum", "-0b1" + strings.Repeat("0", 63), int64(math.MinInt64), "-9223372036854775808"},
		{"octal separators", "0o_7_1", int64(57), "57"},
		{"octal maximum", "0o777777777777777777777", int64(math.MaxInt64), "9223372036854775807"},
		{"octal minimum", "-0o1000000000000000000000", int64(math.MinInt64), "-9223372036854775808"},
		{"dozenal digits", "0z_XE", int64(131), "131"},
		{"lowercase dozenal", "-0Zxe", int64(-131), "-131"},
		{"hexadecimal separators", "+0x_CA_FE", int64(51966), "51966"},
		{"hexadecimal maximum", "0x7fff_ffff_ffff_ffff", int64(math.MaxInt64), "9223372036854775807"},
		{"hexadecimal minimum", "-0x8000_0000_0000_0000", int64(math.MinInt64), "-9223372036854775808"},
		{"explicit hexadecimal", "HEX:2A", int64(42), "42"},
		{"negative explicit hexadecimal", "-hex:2A", int64(-42), "-42"},
		{"fraction", ".625", 0.625, ".625"},
		{"decimal spelling", "+1_0.50", 10.5, "10.50"},
		{"exponent only", "3e4", float64(30000), "3e4"},
		{"exponent spelling", "+3E+04", float64(30000), "3e04"},
		{"negative exponent", "-2.5E-2", -0.025, "-2.5e-2"},
		{"float negative zero", "-0.0", math.Copysign(0, -1), "-0.0"},
		{"zero with tiny exponent", "0e-400", float64(0), "0e-400"},
		{"smallest subnormal", "5e-324", math.SmallestNonzeroFloat64, "5e-324"},
		{"largest float", "1.7976931348623157e308", math.MaxFloat64, "1.7976931348623157e308"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			literal := test.literal

			// 2. Act.
			got, canonical, err := parseNumber(literal)

			// 3. Assert.
			if err != nil {
				t.Fatalf("parseNumber(%q): %v", literal, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("value = %v (%T), want %v (%T)", got, got, test.want, test.want)
			}
			if canonical != test.canonical {
				t.Errorf("canonical = %q, want %q", canonical, test.canonical)
			}
			if value, ok := got.(float64); ok && value == 0 && math.Signbit(value) != math.Signbit(test.want.(float64)) {
				t.Errorf("zero sign = %v, want %v", math.Signbit(value), math.Signbit(test.want.(float64)))
			}
		})
	}
}

func TestParseNumberErrors(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name    string
		literal string
		message string
	}{
		{"decimal overflow", "9223372036854775808", "signed 64-bit range"},
		{"decimal underflow", "-9223372036854775809", "signed 64-bit range"},
		{"binary overflow", "0b1" + strings.Repeat("0", 63), "signed 64-bit range"},
		{"legacy binary overflow", "%1" + strings.Repeat("0", 63), "signed 64-bit range"},
		{"octal overflow", "0o1000000000000000000000", "signed 64-bit range"},
		{"dozenal overflow", "0z" + strings.Repeat("E", 18), "signed 64-bit range"},
		{"hexadecimal overflow", "0x8000000000000000", "signed 64-bit range"},
		{"explicit hexadecimal overflow", "hex:8000000000000000", "signed 64-bit range"},
		{"invalid base digit", "0b102", "invalid"},
		{"missing digits", "0x", "invalid"},
		{"invalid exponent", "1e", "invalid floating-point"},
		{"positive float overflow", "1e309", "too large"},
		{"negative float overflow", "-1e309", "too large"},
		{"positive float underflow", "1e-400", "too small"},
		{"negative float underflow", "-1e-400", "too small"},
		{"below subnormal rounding boundary", "2e-324", "too small"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			literal := test.literal

			// 2. Act.
			value, canonical, err := parseNumber(literal)

			// 3. Assert.
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Fatalf("error = %v, want a message containing %q", err, test.message)
			}
			if value != nil || canonical != "" {
				t.Errorf("failed conversion returned value %#v and canonical %q", value, canonical)
			}
		})
	}
}
