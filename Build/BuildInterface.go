package Build

import (
	"ProjectCI/Configuration"
	"github.com/google/uuid"
	"go/types"
	"time"
)

type BuildInterface interface {
	getIdentifier() uuid.UUID
	getConfiguration() Configuration.Configuration
	getStartDate() time.Time
	getEndDate() time.Time
	getTags() types.Array
	isValid() bool
	isRunning() bool
	getSteps() types.Array
}
