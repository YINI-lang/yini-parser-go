package integration_test

import (
	"path/filepath"
	"reflect"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func TestDecodeJobRunnerFile(t *testing.T) {
	// 1. Assign.
	type job struct {
		Name     string   `yini:"name"`
		Attempts uint8    `yini:"attempts"`
		Labels   []string `yini:"labels"`
	}
	type retry struct {
		Delays []uint16 `yini:"delays"`
		Jitter float64  `yini:"jitter"`
	}
	type pool struct {
		Enabled bool              `yini:"enabled"`
		Limits  map[string]uint16 `yini:"limits"`
		Retry   retry             `yini:"Retry"`
	}
	type config struct {
		Name     string            `yini:"name"`
		Workers  *uint16           `yini:"workers"`
		Queues   []string          `yini:"queues"`
		Jobs     []job             `yini:"jobs"`
		Fallback *string           `yini:"fallback"`
		MaxDelay float64           `yini:"maxDelay"`
		Metadata map[string]string `yini:"metadata"`
		Pool     *pool             `yini:"Pool"`
	}

	path := filepath.Join("testdata", "strict", "job-runner.strict.yini")
	result, err := yini.ParseFile(path, yini.Options{Mode: yini.ModeStrict})
	if err != nil {
		t.Fatalf("ParseFile: %v (diagnostics: %#v)", err, result.Diagnostics)
	}
	workers := uint16(6)
	want := config{
		Name:    "nightly-index",
		Workers: &workers,
		Queues:  []string{"scan", "summarize"},
		Jobs: []job{
			{Name: "index", Attempts: 3, Labels: []string{"local", "fast"}},
			{Name: "summarize", Attempts: 1, Labels: []string{}},
		},
		MaxDelay: 12.5,
		Metadata: map[string]string{"owner": "operations", "region": "east"},
		Pool: &pool{
			Enabled: true,
			Limits:  map[string]uint16{"pending": 64, "running": 6},
			Retry:   retry{Delays: []uint16{1, 3, 8}, Jitter: 0.125},
		},
	}
	t.Run("document", func(t *testing.T) {
		// 1. Assign.
		var whole struct{ JobRunner config }

		// 2. Act.
		err := result.Document.Decode(&whole)

		// 3. Assert.
		if err != nil {
			t.Fatalf("Document.Decode: %v", err)
		}
		if !reflect.DeepEqual(whole.JobRunner, want) {
			t.Fatalf("Document.Decode = %#v, want %#v", whole.JobRunner, want)
		}
	})
	t.Run("section", func(t *testing.T) {
		// 1. Assign.
		fallback := "old fallback"
		section := config{Fallback: &fallback}

		// 2. Act.
		err := result.Document.Sections[0].Decode(&section)

		// 3. Assert.
		if err != nil {
			t.Fatalf("Section.Decode: %v", err)
		}
		if !reflect.DeepEqual(section, want) {
			t.Fatalf("Section.Decode = %#v, want %#v", section, want)
		}
	})
	t.Run("independent copies", func(t *testing.T) {
		// 1. Assign.
		var whole struct{ JobRunner config }
		if err := result.Document.Decode(&whole); err != nil {
			t.Fatalf("Document.Decode: %v", err)
		}
		var section config
		if err := result.Document.Sections[0].Decode(&section); err != nil {
			t.Fatalf("Section.Decode: %v", err)
		}
		// 2. Act.
		section.Jobs[0].Labels[0] = "changed"
		section.Pool.Limits["running"] = 99

		// 3. Assert.
		if !reflect.DeepEqual(whole.JobRunner, want) {
			t.Fatal("separate Decode calls share collection storage")
		}
		t.Run("source unchanged", func(t *testing.T) {
			// 1. Assign.
			var fresh config

			// 2. Act.
			err := result.Document.Sections[0].Decode(&fresh)

			// 3. Assert.
			if err != nil {
				t.Fatalf("Section.Decode after modifying destination: %v", err)
			}
			if !reflect.DeepEqual(fresh, want) {
				t.Fatalf("decoding modified the source: got %#v, want %#v", fresh, want)
			}
		})
	})
}
