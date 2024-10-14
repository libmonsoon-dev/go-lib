//go:build !darwin && !windows

package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

const (
	binName = "notify-send"
)

var (
	binPath   string
	checkErr  error
	checkOnce sync.Once
)

func checkBin() (err error) {
	checkOnce.Do(func() {
		binPath, err = exec.LookPath(binName)
		if err != nil {
			checkErr = fmt.Errorf("could not find %s executable: %w (try installing with your favorite package manager)", binName, err)
			return
		}
	})

	return checkErr
}

// $ notify-send --version
// notify-send 0.7.9
// $ notify-send --help
// Usage:
// notify-send [OPTION…] <SUMMARY> [BODY] - create a notification
//
// Help Options:
// -?, --help                        Show help options
//
// Application Options:
// -u, --urgency=LEVEL               Specifies the urgency level (low, normal, critical).
// -t, --expire-time=TIME            Specifies the timeout in milliseconds at which to expire the notification.
// -a, --app-name=APP_NAME           Specifies the app name for the icon
// -i, --icon=ICON[,ICON...]         Specifies an icon filename or stock icon to display.
// -c, --category=TYPE[,TYPE...]     Specifies the notification category.
// -h, --hint=TYPE:NAME:VALUE        Specifies basic extra data to pass. Valid types are int, double, string and byte.
// -v, --version                     Version of the package.
func (n Notification) newSendCommand(ctx context.Context, summary, body string) (*exec.Cmd, func(), error) {
	err := checkBin()
	if err != nil {
		return nil, nil, err
	}

	var args []string
	if urgency, ok := n.urgencyLevelString(); ok {
		args = append(args, "--urgency", urgency)
	}

	if n.Expire > 0 {
		args = append(args, "--expire-time", strconv.FormatInt(n.Expire.Milliseconds(), 10))
	}

	if n.AppName != "" {
		args = append(args, "--app-name", n.AppName)
	}

	if n.Icon != "" {
		args = append(args, "--icon", n.Icon)
	}

	if len(n.Category) > 0 {
		args = append(args, "--category", strings.Join(n.Category, ","))
	}

	for key, value := range n.Hint {
		var typ, strValue string

		switch value := value.(type) {
		case byte:
			typ, strValue = "byte", fmt.Sprintf("%d", value)
		case int, int8, int16, int32, int64, uint16, uint32, uint64:
			typ, strValue = "int", fmt.Sprintf("%d", value)
		case float32, float64:
			typ, strValue = "double", fmt.Sprintf("%f", value)
		case string:
			typ, strValue = "string", value
		default:
			data, err := json.Marshal(value)
			if err != nil {
				continue
			}

			typ, strValue = "string", string(data)
		}

		args = append(args, "--hint", strings.Join([]string{typ, key, strValue}, ":"))
	}

	return exec.CommandContext(ctx, binPath, append(args, summary, body)...), nil, nil
}

func (n Notification) urgencyLevelString() (string, bool) {
	str, ok := map[UrgencyLevel]string{
		UrgencyLevelLow:      "low",
		UrgencyLevelNormal:   "normal",
		UrgencyLevelCritical: "critical",
	}[n.UrgencyLevel]

	return str, ok
}
