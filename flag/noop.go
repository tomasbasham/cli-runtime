package flag

import (
	goflag "flag"

	"github.com/spf13/pflag"
)

// Noop is a no-op flag value.
type Noop struct{}

var _ goflag.Value = &Noop{}
var _ pflag.Value = &Noop{}

func (n *Noop) String() string {
	return ""
}

func (n *Noop) Set(string) error {
	return nil
}

func (n *Noop) Type() string {
	return "noop"
}
