package notification_test

import (
	"context"
	"os/exec"
	"runtime"
	"testing"

	"github.com/libmonsoon-dev/go-lib/notification"
)

func TestSend(t *testing.T) {
	var binName string
	switch runtime.GOOS {
	case "linux":
		binName = "notify-send"
	case "darwin":
		binName = "terminal-notifier"
	case "windows":
		binName = "PowerShell"
	default:
		t.Skipf("unexpected os %s", runtime.GOOS)
	}

	_, err := exec.LookPath(binName)
	if err != nil {
		t.Skip(err)
	}

	err = notification.Send(context.Background(), "title", "body")
	if err != nil {
		t.Fatal(err)
	}
}
