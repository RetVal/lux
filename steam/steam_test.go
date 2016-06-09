package steam

import (
	"testing"
)

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func TestGetSteamFriends(t *testing.T) {
	f := Friends()
	name := f.GetPersonaName()
	t.Log(name)
	if name == "" {
		t.Error("name was empty")
	}
}
