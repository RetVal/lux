package steam

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	if !Init() {
		fmt.Println("Can't init")
		return
	}
	m.Run()
	Shutdown()
}

func TestGetSteamFriends(t *testing.T) {
	f := Friends()
	name := f.GetPersonaName()
	t.Log(name)
	if name == "" {
		t.Error("name was empty")
	}
}

func TestHTMLSurface(t *testing.T) {
	surface := HTMLSurface()
	t.Log("surface", surface)
	if surface.Pointer == nil {
		t.Error("surface.Pointer is nil")
		return
	}

	if !surface.Init() {
		t.Error("Cannot Init HTMLSurface")
		return
	}

	browser := surface.CreateBrowser("", "")
	t.Log("browser", browser)
	if browser == 0 {
		t.Error("browser == 0")
		return
	}

	surface.LoadURL(browser, "https://www.google.com", "")

	surface.RemoveBrowser(browser)

	if !surface.Shutdown() {
		t.Error("Cannot Shutdown HTMLSurface")
		return
	}
}
