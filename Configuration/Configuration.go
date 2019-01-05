package Configuration

import (
	"github.com/google/uuid"
)

type Configuration struct {
	identifier uuid.UUID
	fileName string
	changed bool
	valid bool
}

func (config *Configuration) getIdentifier() uuid.UUID {
	return config.identifier
}

func (config *Configuration) getIdentifierAsString() string {
	return config.identifier.String()
}

func (config *Configuration) getFileName() string {
	return config.fileName
}

func (config *Configuration) hasChanged() bool {
	return config.changed
}

func (config *Configuration) isValid() bool {
	return config.valid
}
