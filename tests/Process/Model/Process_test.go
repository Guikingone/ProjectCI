package Model

import (
	"ProjectCI/Process/Model"
	"testing"
)

func TestProcessCreationAndRunning(t *testing.T) {
	process := Model.Process{}
	process.New()

	if process.IsRunning() == false {
		t.Errorf("The process isn't running")
	}
}

func TestProcessRestart (t *testing.T) {
	process := Model.Process{}
	process.New()

	process.Pause()

	if process.CanBeRestarted() == false {
		t.Errorf("The process could be restarted as it's paused")
	}

	process.Restart()


	if process.CanBeRestarted() == true {
		t.Errorf("The process shouldn't be restarted as it's running")
	}
}

func TestProcessStop(t *testing.T) {
	process := Model.Process{}
	process.New()

	process.Stop()

	if process.IsRunning() == true {
		t.Errorf("The process must be stopped")
	}

	if process.CanBeRestarted() == true {
		t.Errorf("The process shouldn't be restarted")
	}
}
