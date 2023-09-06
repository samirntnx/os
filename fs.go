package os

import (
	"io"
	"io/fs"
	"os"
	"syscall"
	"time"
)

type any = interface{}

// This file is generated using samirntnx/mockpackage

type OsInterface interface {
	Setenv(key, value string) error
	Getppid() int
	Create(name string) (FileInterface, error)
	UserHomeDir() (string, error)
	Pipe() (r FileInterface, w FileInterface, err error)
	Getuid() int
	SameFile(fi1, fi2 os.FileInfo) bool
	Clearenv()
	Getpid() int
	UserCacheDir() (string, error)
	ReadFile(name string) ([]byte, error)
	Getpagesize() int
	Unsetenv(key string) error
	FindProcess(pid int) (ProcessInterface, error)
	Chmod(name string, mode os.FileMode) error
	Lstat(name string) (os.FileInfo, error)
	Rename(oldpath, newpath string) error
	DirFS(dir string) fs.FS
	WriteFile(name string, data []byte, perm os.FileMode) error
	NewSyscallError(syscall string, err error) error
	IsExist(err error) bool
	OpenFile(name string, flag int, perm os.FileMode) (FileInterface, error)
	MkdirTemp(dir, pattern string) (string, error)
	ReadDir(name string) ([]os.DirEntry, error)
	Symlink(oldname, newname string) error
	IsPathSeparator(c uint8) bool
	Getegid() int
	ExpandEnv(s string) string
	Link(oldname, newname string) error
	Open(name string) (FileInterface, error)
	NewFile(fd uintptr, name string) FileInterface
	Chown(name string, uid, gid int) error
	Exit(code int)
	Environ() []string
	IsNotExist(err error) bool
	IsTimeout(err error) bool
	UserConfigDir() (string, error)
	Lchown(name string, uid, gid int) error
	Getgid() int
	Chdir(dir string) error
	Getenv(key string) string
	Mkdir(name string, perm os.FileMode) error
	TempDir() string
	Geteuid() int
	Getgroups() ([]int, error)
	Stat(name string) (os.FileInfo, error)
	Chtimes(name string, atime time.Time, mtime time.Time) error
	Remove(name string) error
	Readlink(name string) (string, error)
	Getwd() (dir string, err error)
	RemoveAll(path string) error
	Hostname() (name string, err error)
	Expand(s string, mapping func(string) string) string
	LookupEnv(key string) (string, bool)
	IsPermission(err error) bool
	StartProcess(name string, argv []string, attr *os.ProcAttr) (ProcessInterface, error)
	Executable() (string, error)
	Truncate(name string, size int64) error
	MkdirAll(path string, perm os.FileMode) error
	CreateTemp(dir, pattern string) (FileInterface, error)
}

type FileInterface interface {
	Readdir(n int) ([]os.FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	ReadDir(n int) ([]os.DirEntry, error)
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadFrom(r io.Reader) (n int64, err error)
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	WriteString(s string) (n int, err error)
	Chmod(mode os.FileMode) error
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Close() error
	Chown(uid, gid int) error
	Truncate(size int64) error
	Sync() error
	Chdir() error
	Fd() uintptr
	Stat() (os.FileInfo, error)
}

type SyscallErrorInterface interface {
	Error() string
	Unwrap() error
	Timeout() bool
}

type ProcessInterface interface {
	Release() error
	Kill() error
	Wait() (ProcessStateInterface, error)
	Signal(sig os.Signal) error
}

type ProcessStateInterface interface {
	UserTime() time.Duration
	SystemTime() time.Duration
	Exited() bool
	Success() bool
	Sys() any
	SysUsage() any
	Pid() int
	String() string
	ExitCode() int
}

type LinkErrorInterface interface {
	Error() string
	Unwrap() error
}

type OsImpl struct{}

func (recv *OsImpl) Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

func (recv *OsImpl) Exit(code int) {
	os.Exit(code)
}

func (recv *OsImpl) Environ() []string {
	return os.Environ()
}

func (recv *OsImpl) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (recv *OsImpl) IsTimeout(err error) bool {
	return os.IsTimeout(err)
}

func (recv *OsImpl) UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

func (recv *OsImpl) Lchown(name string, uid, gid int) error {
	return os.Lchown(name, uid, gid)
}

func (recv *OsImpl) Getgid() int {
	return os.Getgid()
}

func (recv *OsImpl) Chdir(dir string) error {
	return os.Chdir(dir)
}

func (recv *OsImpl) Getenv(key string) string {
	return os.Getenv(key)
}

func (recv *OsImpl) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func (recv *OsImpl) TempDir() string {
	return os.TempDir()
}

func (recv *OsImpl) Geteuid() int {
	return os.Geteuid()
}

func (recv *OsImpl) Getgroups() ([]int, error) {
	return os.Getgroups()
}

func (recv *OsImpl) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (recv *OsImpl) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func (recv *OsImpl) Remove(name string) error {
	return os.Remove(name)
}

func (recv *OsImpl) Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func (recv *OsImpl) Getwd() (dir string, err error) {
	return os.Getwd()
}

func (recv *OsImpl) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (recv *OsImpl) Hostname() (name string, err error) {
	return os.Hostname()
}

func (recv *OsImpl) Expand(s string, mapping func(string) string) string {
	return os.Expand(s, mapping)
}

func (recv *OsImpl) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (recv *OsImpl) IsPermission(err error) bool {
	return os.IsPermission(err)
}

func (recv *OsImpl) StartProcess(name string, argv []string, attr *os.ProcAttr) (ProcessInterface, error) {
	ret1, err := os.StartProcess(name, argv, attr)
	if ret1 != nil {
		return &ProcessImpl{
			Process: ret1,
		}, err
	}
	return nil, err
}

func (recv *OsImpl) Executable() (string, error) {
	return os.Executable()
}

func (recv *OsImpl) Truncate(name string, size int64) error {
	return os.Truncate(name, size)
}

func (recv *OsImpl) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (recv *OsImpl) CreateTemp(dir, pattern string) (FileInterface, error) {
	return os.CreateTemp(dir, pattern)
}

func (recv *OsImpl) Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func (recv *OsImpl) Getppid() int {
	return os.Getppid()
}

func (recv *OsImpl) Create(name string) (FileInterface, error) {
	return os.Create(name)
}

func (recv *OsImpl) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (recv *OsImpl) Pipe() (r FileInterface, w FileInterface, err error) {
	return os.Pipe()
}

func (recv *OsImpl) Getuid() int {
	return os.Getuid()
}

func (recv *OsImpl) SameFile(fi1, fi2 os.FileInfo) bool {
	return os.SameFile(fi1, fi2)
}

func (recv *OsImpl) Clearenv() {
	os.Clearenv()
}

func (recv *OsImpl) Getpid() int {
	return os.Getpid()
}

func (recv *OsImpl) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

func (recv *OsImpl) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (recv *OsImpl) Getpagesize() int {
	return os.Getpagesize()
}

func (recv *OsImpl) Unsetenv(key string) error {
	return os.Unsetenv(key)
}

func (recv *OsImpl) FindProcess(pid int) (ProcessInterface, error) {
	ret1, err := os.FindProcess(pid)
	if ret1 != nil {
		return &ProcessImpl{
			Process: ret1,
		}, err
	}
	return nil, err
}

func (recv *OsImpl) Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func (recv *OsImpl) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (recv *OsImpl) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (recv *OsImpl) DirFS(dir string) fs.FS {
	return os.DirFS(dir)
}

func (recv *OsImpl) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (recv *OsImpl) NewSyscallError(syscall string, err error) error {
	return os.NewSyscallError(syscall, err)
}

func (recv *OsImpl) IsExist(err error) bool {
	return os.IsExist(err)
}

func (recv *OsImpl) OpenFile(name string, flag int, perm os.FileMode) (FileInterface, error) {
	return os.OpenFile(name, flag, perm)
}

func (recv *OsImpl) MkdirTemp(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

func (recv *OsImpl) ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

func (recv *OsImpl) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

func (recv *OsImpl) IsPathSeparator(c uint8) bool {
	return os.IsPathSeparator(c)
}

func (recv *OsImpl) Getegid() int {
	return os.Getegid()
}

func (recv *OsImpl) ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func (recv *OsImpl) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}

func (recv *OsImpl) Open(name string) (FileInterface, error) {
	return os.Open(name)
}

func (recv *OsImpl) NewFile(fd uintptr, name string) FileInterface {
	return os.NewFile(fd, name)
}

type FileImpl struct {
	*os.File
}

func (recv *FileImpl) Readdir(n int) ([]os.FileInfo, error) {
	return recv.File.Readdir(n)
}

func (recv *FileImpl) Readdirnames(n int) (names []string, err error) {
	return recv.File.Readdirnames(n)
}

func (recv *FileImpl) ReadDir(n int) ([]os.DirEntry, error) {
	return recv.File.ReadDir(n)
}

func (recv *FileImpl) Name() string {
	return recv.File.Name()
}

func (recv *FileImpl) Read(b []byte) (n int, err error) {
	return recv.File.Read(b)
}

func (recv *FileImpl) ReadAt(b []byte, off int64) (n int, err error) {
	return recv.File.ReadAt(b, off)
}

func (recv *FileImpl) ReadFrom(r io.Reader) (n int64, err error) {
	return recv.File.ReadFrom(r)
}

func (recv *FileImpl) Write(b []byte) (n int, err error) {
	return recv.File.Write(b)
}

func (recv *FileImpl) WriteAt(b []byte, off int64) (n int, err error) {
	return recv.File.WriteAt(b, off)
}

func (recv *FileImpl) Seek(offset int64, whence int) (ret int64, err error) {
	return recv.File.Seek(offset, whence)
}

func (recv *FileImpl) WriteString(s string) (n int, err error) {
	return recv.File.WriteString(s)
}

func (recv *FileImpl) Chmod(mode os.FileMode) error {
	return recv.File.Chmod(mode)
}

func (recv *FileImpl) SetDeadline(t time.Time) error {
	return recv.File.SetDeadline(t)
}

func (recv *FileImpl) SetReadDeadline(t time.Time) error {
	return recv.File.SetReadDeadline(t)
}

func (recv *FileImpl) SetWriteDeadline(t time.Time) error {
	return recv.File.SetWriteDeadline(t)
}

func (recv *FileImpl) SyscallConn() (syscall.RawConn, error) {
	return recv.File.SyscallConn()
}

func (recv *FileImpl) Close() error {
	return recv.File.Close()
}

func (recv *FileImpl) Chown(uid, gid int) error {
	return recv.File.Chown(uid, gid)
}

func (recv *FileImpl) Truncate(size int64) error {
	return recv.File.Truncate(size)
}

func (recv *FileImpl) Sync() error {
	return recv.File.Sync()
}

func (recv *FileImpl) Chdir() error {
	return recv.File.Chdir()
}

func (recv *FileImpl) Fd() uintptr {
	return recv.File.Fd()
}

func (recv *FileImpl) Stat() (os.FileInfo, error) {
	return recv.File.Stat()
}

type SyscallErrorImpl struct {
	*os.SyscallError
}

func (recv *SyscallErrorImpl) Error() string {
	return recv.SyscallError.Error()
}

func (recv *SyscallErrorImpl) Unwrap() error {
	return recv.SyscallError.Unwrap()
}

func (recv *SyscallErrorImpl) Timeout() bool {
	return recv.SyscallError.Timeout()
}

type ProcessImpl struct {
	*os.Process
}

func (recv *ProcessImpl) Release() error {
	return recv.Process.Release()
}

func (recv *ProcessImpl) Kill() error {
	return recv.Process.Kill()
}

func (recv *ProcessImpl) Wait() (ProcessStateInterface, error) {
	return recv.Process.Wait()
}

func (recv *ProcessImpl) Signal(sig os.Signal) error {
	return recv.Process.Signal(sig)
}

type ProcessStateImpl struct {
	*os.ProcessState
}

func (recv *ProcessStateImpl) UserTime() time.Duration {
	return recv.ProcessState.UserTime()
}

func (recv *ProcessStateImpl) SystemTime() time.Duration {
	return recv.ProcessState.SystemTime()
}

func (recv *ProcessStateImpl) Exited() bool {
	return recv.ProcessState.Exited()
}

func (recv *ProcessStateImpl) Success() bool {
	return recv.ProcessState.Success()
}

func (recv *ProcessStateImpl) Sys() any {
	return recv.ProcessState.Sys()
}

func (recv *ProcessStateImpl) SysUsage() any {
	return recv.ProcessState.SysUsage()
}

func (recv *ProcessStateImpl) Pid() int {
	return recv.ProcessState.Pid()
}

func (recv *ProcessStateImpl) String() string {
	return recv.ProcessState.String()
}

func (recv *ProcessStateImpl) ExitCode() int {
	return recv.ProcessState.ExitCode()
}

type LinkErrorImpl struct {
	*os.LinkError
}

func (recv *LinkErrorImpl) Error() string {
	return recv.LinkError.Error()
}

func (recv *LinkErrorImpl) Unwrap() error {
	return recv.LinkError.Unwrap()
}

var FS = &OsImpl{}
