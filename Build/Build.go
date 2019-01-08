package Build

import (
	"ProjectCI/Configuration"
	"ProjectCI/Filesystem"
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
	repositoryName string
	repositoryBranch string
	repositoryUrl string
	repositoryLanguage string
	repositoryLogName string
	repositoryLogFilepath string
}

func (build *Build) New() {
	fileSystem := Filesystem.Filesystem{}
	// Create a directory for the actual buid
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

func (build *Build) getRepositoryName() string {
	return build.repositoryName
}

func (build *Build) getRepositoryBranch() string {
	return build.repositoryBranch
}

func (build *Build) getRepositoryUrl() string {
	return build.repositoryUrl
}

func (build *Build) getRepositoryLanguage() string {
	return build.repositoryLanguage
}

func (build *Build) getRepositoryLogName() string {
	return build.repositoryLogName
}

func (build *Build) getRepositoryLogFilePath() string {
	return build.repositoryLogFilepath
}
