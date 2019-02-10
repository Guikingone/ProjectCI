package Model

import "github.com/google/uuid"

type ProcessInterface interface {
	New()
	Pause()
	Restart()
	Stop()
	GetIdentifier() uuid.UUID
	IsRunning() bool
	CanBeRestarted() bool
}
