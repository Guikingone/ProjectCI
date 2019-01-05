package Configuration

import "github.com/google/uuid"

type ConfigurationInterface interface {
	getIdentifier() uuid.UUID
	getIdentifierAsString() string
	getFileName() string
	hasChanged() bool
	isValid() bool
}
