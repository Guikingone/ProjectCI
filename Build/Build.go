package Build

import (
	"ProjectCI/Configuration"
	"github.com/google/uuid"
	"go/types"
	"time"
)

type Build struct {
	identifier uuid.UUID
	configuration Configuration.Configuration
	startDate time.Time
	endDate time.Time
	tags types.Array
	valid bool
	running bool
	steps types.Array
}

func (build *Build) getIdentifier() uuid.UUID {
	return build.identifier
}

func (build *Build) getConfiguration() Configuration.Configuration {
	return build.configuration
}

func (build *Build) getStartDate() time.Time {
	return build.startDate
}

func (build *Build) getEndDate() time.Time {
	return build.endDate
}

func (build *Build) getTags() types.Array {
	return build.tags
}

func (build *Build) isValid() bool {
	return build.valid
}

func (build *Build) isRunning() bool {
	return build.running
}

func (build *Build) getSteps() types.Array {
	return build.steps
}
