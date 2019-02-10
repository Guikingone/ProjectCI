package Model

import "github.com/google/uuid"

type Process struct {
	identifier uuid.UUID
	Running bool
	canRestart bool
}

func (process *Process) New() {
	process.identifier = uuid.UUID{}
	process.Running = true
}

func (process *Process) Pause() {
	if process.IsRunning() {
		process.Running = false
		process.canRestart = true

		return
	}

	process.Running = true
	process.canRestart = true
}

func (process *Process) Restart() {
	if process.IsRunning() {
		return
	}

	process.Running = true
	process.canRestart = false
}

func (process *Process) Stop() {
	if !process.IsRunning() {
		return
	}

	process.Running = false
	process.canRestart = false
}

func (process *Process) GetIdentifier() uuid.UUID {
	return process.identifier
}

func (process *Process) IsRunning() bool {
	return process.Running
}

func (process *Process) CanBeRestarted() bool {
	return process.canRestart
}
