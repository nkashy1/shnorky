package components

import (
	"encoding/json"
	"io"
)

// ComponentSpecification - struct specifying how a component of a simplex data processing flow
// should be built and executed
type ComponentSpecification struct {
	Build BuildSpecification `json:"build"`
	Run   RunSpecification   `json:"run"`
}

// BuildSpecification - struct specifying how a component of a simplex data processing flow should
// be built; all paths are assumed to be paths relative to the component path (i.e. the directory
// containing the implementation of the component)
type BuildSpecification struct {
	// Path to Dockerfile to be used to build the component
	Dockerfile string `json:"Dockerfile"`

	// Context that should be provided at build time
	Context string `json:"context"`
}

// RunSpecification - struct specifying how a component of a simplex data processing flow should be
// executed
type RunSpecification struct {
	// Mapping of environment variable names to values to be set in component container at runtime
	Env map[string]string `json:"env"`

	// Command to be invoked when starting component container at runtime
	Cmd []string `json:"cmd"`

	// Mountpoint specify paths inside each container (for this component) that can accept data
	Mountpoints []MountSpecification `json:"mountpoints"`
}

// MountSpecification - specifies a mount point within a simplex component, how it should be mounted
// on the container side, and whether or not it is required to be mounted at runtime
type MountSpecification struct {
	Mountpoint string `json:"mountpoint"`
	ReadOnly   bool   `json:"read_only"`
	Required   bool   `json:"required"`
}

// ReadSingleSpecification reads a single ComponentSpecification JSON document and returns the
// corresponding ComponentSpecification struct. It returns an error if there was an issue parsing
// the specification into the struct.
func ReadSingleSpecification(reader io.Reader) (ComponentSpecification, error) {
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	var specification ComponentSpecification
	err := dec.Decode(&specification)
	return specification, err
}
