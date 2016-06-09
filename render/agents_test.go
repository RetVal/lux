package lux

import (
	"errors"
	"testing"
	"time"
)

func TestNewAgentManager(t *testing.T) {
	NewAgentManager()
}

func TestNewAgent(t *testing.T) {
	am := NewAgentManager()
	ping := make(chan struct{}, 3)
	am.NewAgent(func() bool {
		ping <- struct{}{}
		return true
	})
	am.Tick()
	am.Tick()
	am.Tick()
	count := 0
	select {
	case <-ping:
		count++
		if count == 3 {
			break
		}
	case <-time.After(100 * time.Millisecond):
		t.Error(errors.New("couldnt retreive pings"))
	}
}

func TestAgentSeppuku(t *testing.T) {
	am := NewAgentManager()
	ping := make(chan struct{}, 1)
	agent := am.NewAgent(func() bool {
		ping <- struct{}{}
		return true
	})
	agent.Seppuku()
	am.Tick()
	select {
	case <-ping:
		t.Error(errors.New("there the agent should not have executed"))
	case <-time.After(100 * time.Millisecond):
		break
	}

}

func TestAgentNaturalDeath(t *testing.T) {
	am := NewAgentManager()
	countdown := 2
	am.NewAgent(func() bool {
		countdown--
		return countdown != 0
	})
	am.Tick()
	am.Tick()
	if am.AgentCount() != 0 {
		t.Error(errors.New("agent still alive"))
	}
}

func TestAgentSleep(t *testing.T) {
	/*am := NewAgentManager()
	var count int
	a := am.NewAgent(func() bool {
		count++
		return true
	})
	am.Tick()
	a.Sleep()
	am.Tick()
	a.Awake()
	am.Tick()

	if count != 2 {
		t.Errorf("count was not 2")
	}*/
}
