package exec

import (
	"context"
	"os/exec"
)

// TODO: add others exported symbols

type Cmd = exec.Cmd

func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	return exec.CommandContext(ctx, name, arg...)
}

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
