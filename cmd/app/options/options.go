package options

import (
	"github.com/spf13/pflag"
)

// ServerRunOptions runs a kubernetes api server.
type ServerRunOptions struct {
	ClusterSet []*string
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters
func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{}
	return &s
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
}
