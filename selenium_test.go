package selenium

import (
	"testing"
)

func TestStatus(t *testing.T) {
	wd, err := New(nil, "", nil)
	if err != nil {
		t.Error(err)
	}

	status, err := wd.Status()
	if err != nil {
		t.Error(err)
	}

	if len(status.OS.Name) == 0 {
		t.Error("No OS")
	}
}

func TestNewSession(t *testing.T) {
	caps := &Capabilities {
		"browserName": "firefox",
	}

	wd := &WebDriver{Capabilities: caps, Executor: DEFAULT_EXECUTOR}
	sid, err := wd.NewSession()
	defer wd.Quit()

	if err != nil {
		t.Errorf("error in new session - %s", err)
	}

	if len(sid) == 0 {
		t.Error("Empty session id")
	}

	if wd.SessionId != sid {
		t.Error("Session id mismatch")
	}
}


func TestURL(t *testing.T) {
	params := Params{SessionId: "17"}
	url, err := CMD_QUIT.URL(&params)

	if err != nil {
		t.Error(err)
	}

	if url != "/session/17" {
		t.Errorf("Bad url: %s", url)
	}
}
