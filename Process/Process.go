package Process

import "github.com/google/uuid"

type Process struct {
	identifier uuid.UUID
	Running bool
}

func (process *Process) New() {
	process.identifier = uuid.UUID{}

}

func (process *Process) Pause() {
	if process.isRunning() {
		return
	}

}

func (process *Process) Stop() {
	if process.isRunning() {
		return
	}


}

func (process *Process) getIdentifier() uuid.UUID {
	return process.identifier
}

func (process *Process) isRunning() bool {
	return process.Running
}
