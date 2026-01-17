package printer_test

import (
	"bytes"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"

	"github.com/tomasbasham/cli-runtime/printer"
)

func TestYAMLPrinter_Print(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		documentStart bool
		input         any
		want          string
		wantErr       bool
	}{
		"simple map without document start": {
			input: map[string]any{"key": "value", "number": 42},
			want:  "key: value\nnumber: 42\n",
		},
		"simple map with document start": {
			documentStart: true,
			input:         map[string]any{"key": "value", "number": 42},
			want:          "---\nkey: value\nnumber: 42\n",
		},
		"nested structure": {
			input: map[string]any{"outer": map[string]any{"inner": "value"}},
			want:  "outer:\n  inner: value\n",
		},
		"array": {
			input: []string{"one", "two", "three"},
			want:  "- one\n- two\n- three\n",
		},
		"empty map": {
			input: map[string]any{},
			want:  "{}\n",
		},
		"struct input": {
			input: struct {
				Name string `yaml:"name"`
				Age  int    `yaml:"age"`
			}{
				Name: "Alice",
				Age:  30,
			},
			want: "name: Alice\nage: 30\n",
		},
		"invalid input (channel)": {
			input:   make(chan int),
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Set up YAMLPrinter.
			p := &printer.YAMLPrinter{DocumentStart: tt.documentStart}

			var buf bytes.Buffer
			err := p.Print(&buf, tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if !tt.wantErr {
				var got, want any
				if err := yaml.Unmarshal(buf.Bytes(), &got); err != nil {
					t.Fatalf("invalid YAML output: %v", err)
				}
				if err := yaml.Unmarshal([]byte(tt.want), &want); err != nil {
					t.Fatalf("invalid test fixture: %v", err)
				}

				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
