package printer_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tomasbasham/cli-runtime/printer"
)

func TestJSONPrinter_Print(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		indent  bool
		input   any
		want    string
		wantErr bool
	}{
		"simple map without indentation": {
			input: map[string]any{"key": "value", "number": 42},
			want:  `{"key":"value","number":42}`,
		},
		"simple map with indentation": {
			indent: true,
			input:  map[string]any{"key": "value", "number": 42},
			want: `{
  "key": "value",
	"number": 42
}`,
		},
		"nested structure with indentation": {
			indent: true,
			input:  map[string]any{"outer": map[string]any{"inner": "value"}},
			want: `{
  "outer": {
    "inner": "value"
  }
}`,
		},
		"array without indentation": {
			input: []string{"one", "two", "three"},
			want:  `["one","two","three"]`,
		},
		"empty map": {
			input: map[string]any{},
			want:  `{}`,
		},
		"struct input": {
			input: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{
				Name: "Alice",
				Age:  30,
			},
			want: `{"name":"Alice","age":30}`,
		},
		"invalid input (channel)": {
			input:   make(chan int),
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Set up JSONPrinter.
			p := &printer.JSONPrinter{Indent: tt.indent}

			var buf bytes.Buffer
			err := p.Print(&buf, tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if !tt.wantErr {
				var got, want any
				if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
					t.Fatalf("invalid JSON output: %v", err)
				}
				if err := json.Unmarshal([]byte(tt.want), &want); err != nil {
					t.Fatalf("invalid test fixture: %v", err)
				}

				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
