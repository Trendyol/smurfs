package process

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/trendyol/smurfs/go/host/pkg/plugin"
)

type Executor interface {
	// Run executes the command with the given arguments
	Run(ctx context.Context, receipt *plugin.Receipt, args ...string) error
}

type executor struct {
	hostAddress *string
}

func NewExec(hostAddress *string) Executor {
	return &executor{
		hostAddress: hostAddress,
	}
}

func (e *executor) Run(ctx context.Context, receipt *plugin.Receipt, args ...string) error {
	executablePath := receipt.Executable.Executable.Entrypoint

	hostArgs := []string{
		"--smurf-host-address", *e.hostAddress,
	}

	cmd := exec.CommandContext(ctx, executablePath, append(args, hostArgs...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	fmt.Println(err)
	return err
}
