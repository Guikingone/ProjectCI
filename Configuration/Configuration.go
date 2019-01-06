package Configuration

import (
	"go/types"
)

type Configuration struct {
	Version string      `json:"version"`
	Steps   types.Array `json:"steps"`
	Caches  types.Array `json:"caches"`
}

func (config *Configuration) getVersion() string {
	return config.Version
}

func (config *Configuration) getSteps() types.Array {
	return config.Steps
}

func (config *Configuration) getCaches() types.Array {
	return config.Caches
}
