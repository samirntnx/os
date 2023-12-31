// Code generated by MockGen. DO NOT EDIT.
// Source: exec/exec.go

// Package exec is a generated GoMock package.
package exec

import (
	context "context"
	io "io"
	exec "os/exec"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockany is a mock of any interface.
type Mockany struct {
	ctrl     *gomock.Controller
	recorder *MockanyMockRecorder
}

// MockanyMockRecorder is the mock recorder for Mockany.
type MockanyMockRecorder struct {
	mock *Mockany
}

// NewMockany creates a new mock instance.
func NewMockany(ctrl *gomock.Controller) *Mockany {
	mock := &Mockany{ctrl: ctrl}
	mock.recorder = &MockanyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockany) EXPECT() *MockanyMockRecorder {
	return m.recorder
}

// MockExecInterface is a mock of ExecInterface interface.
type MockExecInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExecInterfaceMockRecorder
}

// MockExecInterfaceMockRecorder is the mock recorder for MockExecInterface.
type MockExecInterfaceMockRecorder struct {
	mock *MockExecInterface
}

// NewMockExecInterface creates a new mock instance.
func NewMockExecInterface(ctrl *gomock.Controller) *MockExecInterface {
	mock := &MockExecInterface{ctrl: ctrl}
	mock.recorder = &MockExecInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecInterface) EXPECT() *MockExecInterfaceMockRecorder {
	return m.recorder
}

// Command mocks base method.
func (m *MockExecInterface) Command(name, arg string) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", name, arg)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// Command indicates an expected call of Command.
func (mr *MockExecInterfaceMockRecorder) Command(name, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockExecInterface)(nil).Command), name, arg)
}

// CommandContext mocks base method.
func (m *MockExecInterface) CommandContext(ctx context.Context, name, arg string) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandContext", ctx, name, arg)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// CommandContext indicates an expected call of CommandContext.
func (mr *MockExecInterfaceMockRecorder) CommandContext(ctx, name, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandContext", reflect.TypeOf((*MockExecInterface)(nil).CommandContext), ctx, name, arg)
}

// LookPath mocks base method.
func (m *MockExecInterface) LookPath(file string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookPath", file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookPath indicates an expected call of LookPath.
func (mr *MockExecInterfaceMockRecorder) LookPath(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookPath", reflect.TypeOf((*MockExecInterface)(nil).LookPath), file)
}

// MockCmdInterface is a mock of CmdInterface interface.
type MockCmdInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCmdInterfaceMockRecorder
}

// MockCmdInterfaceMockRecorder is the mock recorder for MockCmdInterface.
type MockCmdInterfaceMockRecorder struct {
	mock *MockCmdInterface
}

// NewMockCmdInterface creates a new mock instance.
func NewMockCmdInterface(ctrl *gomock.Controller) *MockCmdInterface {
	mock := &MockCmdInterface{ctrl: ctrl}
	mock.recorder = &MockCmdInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCmdInterface) EXPECT() *MockCmdInterfaceMockRecorder {
	return m.recorder
}

// CombinedOutput mocks base method.
func (m *MockCmdInterface) CombinedOutput() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CombinedOutput")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CombinedOutput indicates an expected call of CombinedOutput.
func (mr *MockCmdInterfaceMockRecorder) CombinedOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CombinedOutput", reflect.TypeOf((*MockCmdInterface)(nil).CombinedOutput))
}

// Environ mocks base method.
func (m *MockCmdInterface) Environ() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environ")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Environ indicates an expected call of Environ.
func (mr *MockCmdInterfaceMockRecorder) Environ() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environ", reflect.TypeOf((*MockCmdInterface)(nil).Environ))
}

// Output mocks base method.
func (m *MockCmdInterface) Output() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output.
func (mr *MockCmdInterfaceMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockCmdInterface)(nil).Output))
}

// Run mocks base method.
func (m *MockCmdInterface) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockCmdInterfaceMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCmdInterface)(nil).Run))
}

// Start mocks base method.
func (m *MockCmdInterface) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockCmdInterfaceMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCmdInterface)(nil).Start))
}

// StderrPipe mocks base method.
func (m *MockCmdInterface) StderrPipe() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StderrPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StderrPipe indicates an expected call of StderrPipe.
func (mr *MockCmdInterfaceMockRecorder) StderrPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StderrPipe", reflect.TypeOf((*MockCmdInterface)(nil).StderrPipe))
}

// StdinPipe mocks base method.
func (m *MockCmdInterface) StdinPipe() (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdinPipe")
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StdinPipe indicates an expected call of StdinPipe.
func (mr *MockCmdInterfaceMockRecorder) StdinPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdinPipe", reflect.TypeOf((*MockCmdInterface)(nil).StdinPipe))
}

// StdoutPipe mocks base method.
func (m *MockCmdInterface) StdoutPipe() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdoutPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StdoutPipe indicates an expected call of StdoutPipe.
func (mr *MockCmdInterfaceMockRecorder) StdoutPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdoutPipe", reflect.TypeOf((*MockCmdInterface)(nil).StdoutPipe))
}

// String mocks base method.
func (m *MockCmdInterface) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockCmdInterfaceMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockCmdInterface)(nil).String))
}

// Wait mocks base method.
func (m *MockCmdInterface) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockCmdInterfaceMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockCmdInterface)(nil).Wait))
}

// MockExitErrorInterface is a mock of ExitErrorInterface interface.
type MockExitErrorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExitErrorInterfaceMockRecorder
}

// MockExitErrorInterfaceMockRecorder is the mock recorder for MockExitErrorInterface.
type MockExitErrorInterfaceMockRecorder struct {
	mock *MockExitErrorInterface
}

// NewMockExitErrorInterface creates a new mock instance.
func NewMockExitErrorInterface(ctrl *gomock.Controller) *MockExitErrorInterface {
	mock := &MockExitErrorInterface{ctrl: ctrl}
	mock.recorder = &MockExitErrorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExitErrorInterface) EXPECT() *MockExitErrorInterfaceMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockExitErrorInterface) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockExitErrorInterfaceMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockExitErrorInterface)(nil).Error))
}

// MockErrorInterface is a mock of ErrorInterface interface.
type MockErrorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockErrorInterfaceMockRecorder
}

// MockErrorInterfaceMockRecorder is the mock recorder for MockErrorInterface.
type MockErrorInterfaceMockRecorder struct {
	mock *MockErrorInterface
}

// NewMockErrorInterface creates a new mock instance.
func NewMockErrorInterface(ctrl *gomock.Controller) *MockErrorInterface {
	mock := &MockErrorInterface{ctrl: ctrl}
	mock.recorder = &MockErrorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockErrorInterface) EXPECT() *MockErrorInterfaceMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockErrorInterface) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockErrorInterfaceMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockErrorInterface)(nil).Error))
}

// Unwrap mocks base method.
func (m *MockErrorInterface) Unwrap() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unwrap")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unwrap indicates an expected call of Unwrap.
func (mr *MockErrorInterfaceMockRecorder) Unwrap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unwrap", reflect.TypeOf((*MockErrorInterface)(nil).Unwrap))
}
