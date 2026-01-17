package templates_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tomasbasham/cli-runtime/templates"
)

func TestLongDesc(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		template string
		want     string
	}{
		"empty input produces empty output": {
			template: "",
			want:     "",
		},
		"single line text is preserved as is": {
			template: "Some text",
			want:     "Some text",
		},
		"consecutive new lines are combined into a single paragraph": {
			template: "Line1\nLine2",
			want:     "Line1 Line2",
		},
		"two paragraphs": {
			template: "Line1\n\nLine2",
			want:     "Line1\n\nLine2",
		},
		"leading and trailing spaces are stripped (single line)": {
			template: "\t  \nThe text line  \n  \t",
			want:     "The text line",
		},
		"leading and trailing spaces are stripped (multi line)": {
			template: "\t  \nLine1\nLine2  \n  \t",
			want:     "Line1 Line2",
		},
		"list items with order": {
			template: "Title\n\n1. First item\n2. Second item\n\nSome text",
			want:     "Title\n\n  1. First item\n  2. Second item\n\nSome text",
		},
		"list items without order": {
			template: "\t\t\t\t\tDescriptions.\n\n * Item.\n * Item2.",
			want:     "Descriptions.\n\n  * Item.\n  * Item2.",
		},
		"with code block": {
			template: "Line1\n\n\tif(true) {\n\t  print('hello')\n\t}\n\nLine2",
			want:     "Line1\n\n        if(true) {\n          print('hello')\n        }\n        \nLine2",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := templates.LongDesc(templates.LongDesc(tt.template))
			if diff := diffLines(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func diffLines(a, b string) string {
	aLines := strings.Split(a, "\n")
	bLines := strings.Split(b, "\n")
	return cmp.Diff(aLines, bLines)
}
