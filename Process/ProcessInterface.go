package Process

import "github.com/google/uuid"

type ProcessInterface interface {
	New()
	Pause()
	Stop()
	getIdentifier() uuid.UUID
	isRunning() bool
}
