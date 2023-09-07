package exec

import (
	"context"
	"io"
	"os/exec"
)

type ExecInterface interface {
	Command(name string, arg string) CmdInterface
	CommandContext(ctx context.Context, name string, arg string) CmdInterface
	LookPath(file string) (string, error)
}

type ErrorInterface interface {
	Error() string
	Unwrap() error
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

type ExecImpl struct{}

func (recv *ExecImpl) Command(name string, arg string) CmdInterface {
	return exec.Command(name, arg)
}

func (recv *ExecImpl) CommandContext(ctx context.Context, name string, arg string) CmdInterface {
	return exec.CommandContext(ctx, name, arg)
}

func (recv *ExecImpl) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

type CmdImpl struct {
	*exec.Cmd
}

func (recv *CmdImpl) String() string {
	return recv.Cmd.String()
}

func (recv *CmdImpl) Run() error {
	return recv.Cmd.Run()
}

func (recv *CmdImpl) Start() error {
	return recv.Cmd.Start()
}

func (recv *CmdImpl) Wait() error {
	return recv.Cmd.Wait()
}

func (recv *CmdImpl) Output() ([]byte, error) {
	return recv.Cmd.Output()
}

func (recv *CmdImpl) CombinedOutput() ([]byte, error) {
	return recv.Cmd.CombinedOutput()
}

func (recv *CmdImpl) StdinPipe() (io.WriteCloser, error) {
	return recv.Cmd.StdinPipe()
}

func (recv *CmdImpl) StdoutPipe() (io.ReadCloser, error) {
	return recv.Cmd.StdoutPipe()
}

func (recv *CmdImpl) StderrPipe() (io.ReadCloser, error) {
	return recv.Cmd.StderrPipe()
}

func (recv *CmdImpl) Environ() []string {
	return recv.Cmd.Environ()
}

var EXEC = &ExecImpl{}
