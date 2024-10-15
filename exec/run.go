package exec

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func Run(ctx context.Context, name string, args ...string) (stdout []byte, stderr []byte, err error) {
	cmd := exec.CommandContext(ctx, name, args...)
	return RunCommand(cmd)
}

func RunCommand(cmd *Cmd) (stdout []byte, stderr []byte, err error) {
	if cmd.Stdout != nil {
		return nil, nil, fmt.Errorf("exec: Stdout already set")
	}
	if cmd.Stderr != nil {
		return nil, nil, fmt.Errorf("exec: Stderr already set")
	}
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err = cmd.Run()
	if err != nil {
		return nil, nil, &Error{cmd, out.Bytes(), errOut.Bytes(), err}
	}

	return out.Bytes(), errOut.Bytes(), nil
}

var _ error = (*Error)(nil)

type Error struct {
	Cmd    *Cmd
	Stdout []byte
	Stderr []byte
	err    error
}

func (e *Error) Error() string {
	cmd := e.Command()
	output := e.Output()
	var err error
	if len(output) > 0 {
		err = fmt.Errorf("execute %s:\n%s\n%w", cmd, string(output), e.err)
	} else {
		err = fmt.Errorf("execute %s: %w", cmd, e.err)
	}

	return err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Output() []byte {
	if len(e.Stdout)+len(e.Stderr) == 0 {
		return nil
	}

	return bytes.Join([][]byte{[]byte("stdout:"), e.Stdout, []byte("stderr:"), e.Stderr}, []byte("\n"))
}

func (e *Error) Command() string {
	args := quoteArgs(e.Cmd.Args)
	return strings.Join(args, " ")
}

func quoteArgs(args []string) []string {
	result := make([]string, len(args))

	for i := range args {
		if strings.Contains(args[i], " ") { //TODO: handle other cases
			result[i] = strconv.Quote(args[i])
		} else {
			result[i] = args[i]
		}
	}

	return result
}
