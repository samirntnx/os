package exec

import (
	context "context"
	io "io"
	"os/exec"
)

type any = interface{}

type ExecInterface interface {
	Command(name string, arg string) *exec.Cmd
	CommandContext(ctx context.Context, name string, arg string) *exec.Cmd
	LookPath(file string) (string, error)
}

type CmdInterface interface {
	String() string
	Run() error
	Start() error
	Wait() error
	Output() ([]byte, error)
	CombinedOutput() ([]byte, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Environ() []string
}

type ExitErrorInterface interface {
	Error() string
}

type ErrorInterface interface {
	Error() string
	Unwrap() error
}

// Test override the below var to provide mockgen mock of package interface
var VarExecMock ExecInterface = nil

func Command(name string, arg string) *exec.Cmd {
	if VarExecMock == nil {
		return VarExecMock.Command(name, arg)
	}
	return exec.Command(name, arg)
}

func CommandContext(ctx context.Context, name string, arg string) *exec.Cmd {
	if VarExecMock == nil {
		return VarExecMock.CommandContext(ctx, name, arg)
	}
	return exec.CommandContext(ctx, name, arg)
}

func LookPath(file string) (string, error) {
	if VarExecMock == nil {
		return VarExecMock.LookPath(file)
	}
	return exec.LookPath(file)
}
