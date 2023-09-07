package os

import (
	io "io"
	fs "io/fs"
	os "os"
	syscall "syscall"
	time "time"
)

type any = interface{}

type OsInterface interface {
	Lchown(name string, uid, gid int) error
	ExpandEnv(s string) string
	Setenv(key, value string) error
	NewSyscallError(syscall string, err error) error
	IsPermission(err error) bool
	StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)
	UserConfigDir() (string, error)
	NewFile(fd uintptr, name string) *os.File
	Executable() (string, error)
	Rename(oldpath, newpath string) error
	Chtimes(name string, atime time.Time, mtime time.Time) error
	Symlink(oldname, newname string) error
	Getenv(key string) string
	Create(name string) (*os.File, error)
	IsPathSeparator(c uint8) bool
	Lstat(name string) (os.FileInfo, error)
	IsTimeout(err error) bool
	UserCacheDir() (string, error)
	Truncate(name string, size int64) error
	Getgid() int
	Expand(s string, mapping func(string) string) string
	Readlink(name string) (string, error)
	MkdirTemp(dir, pattern string) (string, error)
	Getpagesize() int
	LookupEnv(key string) (string, bool)
	Clearenv()
	RemoveAll(path string) error
	Getuid() int
	Hostname() (name string, err error)
	Getegid() int
	Unsetenv(key string) error
	WriteFile(name string, data []byte, perm os.FileMode) error
	Chown(name string, uid, gid int) error
	MkdirAll(path string, perm os.FileMode) error
	CreateTemp(dir, pattern string) (*os.File, error)
	ReadDir(name string) ([]os.DirEntry, error)
	Open(name string) (*os.File, error)
	DirFS(dir string) fs.FS
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	Chmod(name string, mode os.FileMode) error
	Getwd() (dir string, err error)
	Pipe() (r *os.File, w *os.File, err error)
	FindProcess(pid int) (*os.Process, error)
	Exit(code int)
	Mkdir(name string, perm os.FileMode) error
	Chdir(dir string) error
	SameFile(fi1, fi2 os.FileInfo) bool
	Getpid() int
	TempDir() string
	Remove(name string) error
	Geteuid() int
	Getgroups() ([]int, error)
	IsExist(err error) bool
	Environ() []string
	IsNotExist(err error) bool
	Getppid() int
	UserHomeDir() (string, error)
	ReadFile(name string) ([]byte, error)
	Link(oldname, newname string) error
	Stat(name string) (os.FileInfo, error)
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
	Wait() (*os.ProcessState, error)
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

// Test override the below var to provide mockgen mock of package interface
var VarOsMock OsInterface = nil

func IsExist(err error) bool {
	if VarOsMock == nil {
		return VarOsMock.IsExist(err)
	}
	return os.IsExist(err)
}

func ReadFile(name string) ([]byte, error) {
	if VarOsMock == nil {
		return VarOsMock.ReadFile(name)
	}
	return os.ReadFile(name)
}

func Link(oldname, newname string) error {
	if VarOsMock == nil {
		return VarOsMock.Link(oldname, newname)
	}
	return os.Link(oldname, newname)
}

func Stat(name string) (os.FileInfo, error) {
	if VarOsMock == nil {
		return VarOsMock.Stat(name)
	}
	return os.Stat(name)
}

func Environ() []string {
	if VarOsMock == nil {
		return VarOsMock.Environ()
	}
	return os.Environ()
}

func IsNotExist(err error) bool {
	if VarOsMock == nil {
		return VarOsMock.IsNotExist(err)
	}
	return os.IsNotExist(err)
}

func Getppid() int {
	if VarOsMock == nil {
		return VarOsMock.Getppid()
	}
	return os.Getppid()
}

func UserHomeDir() (string, error) {
	if VarOsMock == nil {
		return VarOsMock.UserHomeDir()
	}
	return os.UserHomeDir()
}

func Lchown(name string, uid, gid int) error {
	if VarOsMock == nil {
		return VarOsMock.Lchown(name, uid, gid)
	}
	return os.Lchown(name, uid, gid)
}

func StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error) {
	if VarOsMock == nil {
		return VarOsMock.StartProcess(name, argv, attr)
	}
	return os.StartProcess(name, argv, attr)
}

func UserConfigDir() (string, error) {
	if VarOsMock == nil {
		return VarOsMock.UserConfigDir()
	}
	return os.UserConfigDir()
}

func NewFile(fd uintptr, name string) *os.File {
	if VarOsMock == nil {
		return VarOsMock.NewFile(fd, name)
	}
	return os.NewFile(fd, name)
}

func ExpandEnv(s string) string {
	if VarOsMock == nil {
		return VarOsMock.ExpandEnv(s)
	}
	return os.ExpandEnv(s)
}

func Setenv(key, value string) error {
	if VarOsMock == nil {
		return VarOsMock.Setenv(key, value)
	}
	return os.Setenv(key, value)
}

func NewSyscallError(syscall string, err error) error {
	if VarOsMock == nil {
		return VarOsMock.NewSyscallError(syscall, err)
	}
	return os.NewSyscallError(syscall, err)
}

func IsPermission(err error) bool {
	if VarOsMock == nil {
		return VarOsMock.IsPermission(err)
	}
	return os.IsPermission(err)
}

func Executable() (string, error) {
	if VarOsMock == nil {
		return VarOsMock.Executable()
	}
	return os.Executable()
}

func Rename(oldpath, newpath string) error {
	if VarOsMock == nil {
		return VarOsMock.Rename(oldpath, newpath)
	}
	return os.Rename(oldpath, newpath)
}

func Chtimes(name string, atime time.Time, mtime time.Time) error {
	if VarOsMock == nil {
		return VarOsMock.Chtimes(name, atime, mtime)
	}
	return os.Chtimes(name, atime, mtime)
}

func Symlink(oldname, newname string) error {
	if VarOsMock == nil {
		return VarOsMock.Symlink(oldname, newname)
	}
	return os.Symlink(oldname, newname)
}

func Getenv(key string) string {
	if VarOsMock == nil {
		return VarOsMock.Getenv(key)
	}
	return os.Getenv(key)
}

func Create(name string) (*os.File, error) {
	if VarOsMock == nil {
		return VarOsMock.Create(name)
	}
	return os.Create(name)
}

func IsPathSeparator(c uint8) bool {
	if VarOsMock == nil {
		return VarOsMock.IsPathSeparator(c)
	}
	return os.IsPathSeparator(c)
}

func Lstat(name string) (os.FileInfo, error) {
	if VarOsMock == nil {
		return VarOsMock.Lstat(name)
	}
	return os.Lstat(name)
}

func IsTimeout(err error) bool {
	if VarOsMock == nil {
		return VarOsMock.IsTimeout(err)
	}
	return os.IsTimeout(err)
}

func UserCacheDir() (string, error) {
	if VarOsMock == nil {
		return VarOsMock.UserCacheDir()
	}
	return os.UserCacheDir()
}

func Truncate(name string, size int64) error {
	if VarOsMock == nil {
		return VarOsMock.Truncate(name, size)
	}
	return os.Truncate(name, size)
}

func Getgid() int {
	if VarOsMock == nil {
		return VarOsMock.Getgid()
	}
	return os.Getgid()
}

func Expand(s string, mapping func(string) string) string {
	if VarOsMock == nil {
		return VarOsMock.Expand(s, mapping)
	}
	return os.Expand(s, mapping)
}

func Readlink(name string) (string, error) {
	if VarOsMock == nil {
		return VarOsMock.Readlink(name)
	}
	return os.Readlink(name)
}

func MkdirTemp(dir, pattern string) (string, error) {
	if VarOsMock == nil {
		return VarOsMock.MkdirTemp(dir, pattern)
	}
	return os.MkdirTemp(dir, pattern)
}

func Getpagesize() int {
	if VarOsMock == nil {
		return VarOsMock.Getpagesize()
	}
	return os.Getpagesize()
}

func Hostname() (name string, err error) {
	if VarOsMock == nil {
		return VarOsMock.Hostname()
	}
	return os.Hostname()
}

func LookupEnv(key string) (string, bool) {
	if VarOsMock == nil {
		return VarOsMock.LookupEnv(key)
	}
	return os.LookupEnv(key)
}

func Clearenv() {
	if VarOsMock == nil {
		VarOsMock.Clearenv()
	}
	os.Clearenv()
}

func RemoveAll(path string) error {
	if VarOsMock == nil {
		return VarOsMock.RemoveAll(path)
	}
	return os.RemoveAll(path)
}

func Getuid() int {
	if VarOsMock == nil {
		return VarOsMock.Getuid()
	}
	return os.Getuid()
}

func Getegid() int {
	if VarOsMock == nil {
		return VarOsMock.Getegid()
	}
	return os.Getegid()
}

func CreateTemp(dir, pattern string) (*os.File, error) {
	if VarOsMock == nil {
		return VarOsMock.CreateTemp(dir, pattern)
	}
	return os.CreateTemp(dir, pattern)
}

func Unsetenv(key string) error {
	if VarOsMock == nil {
		return VarOsMock.Unsetenv(key)
	}
	return os.Unsetenv(key)
}

func WriteFile(name string, data []byte, perm os.FileMode) error {
	if VarOsMock == nil {
		return VarOsMock.WriteFile(name, data, perm)
	}
	return os.WriteFile(name, data, perm)
}

func Chown(name string, uid, gid int) error {
	if VarOsMock == nil {
		return VarOsMock.Chown(name, uid, gid)
	}
	return os.Chown(name, uid, gid)
}

func MkdirAll(path string, perm os.FileMode) error {
	if VarOsMock == nil {
		return VarOsMock.MkdirAll(path, perm)
	}
	return os.MkdirAll(path, perm)
}

func ReadDir(name string) ([]os.DirEntry, error) {
	if VarOsMock == nil {
		return VarOsMock.ReadDir(name)
	}
	return os.ReadDir(name)
}

func Open(name string) (*os.File, error) {
	if VarOsMock == nil {
		return VarOsMock.Open(name)
	}
	return os.Open(name)
}

func DirFS(dir string) fs.FS {
	if VarOsMock == nil {
		return VarOsMock.DirFS(dir)
	}
	return os.DirFS(dir)
}

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if VarOsMock == nil {
		return VarOsMock.OpenFile(name, flag, perm)
	}
	return os.OpenFile(name, flag, perm)
}

func Chmod(name string, mode os.FileMode) error {
	if VarOsMock == nil {
		return VarOsMock.Chmod(name, mode)
	}
	return os.Chmod(name, mode)
}

func Getwd() (dir string, err error) {
	if VarOsMock == nil {
		return VarOsMock.Getwd()
	}
	return os.Getwd()
}

func Pipe() (r *os.File, w *os.File, err error) {
	if VarOsMock == nil {
		return VarOsMock.Pipe()
	}
	return os.Pipe()
}

func FindProcess(pid int) (*os.Process, error) {
	if VarOsMock == nil {
		return VarOsMock.FindProcess(pid)
	}
	return os.FindProcess(pid)
}

func Exit(code int) {
	if VarOsMock == nil {
		VarOsMock.Exit(code)
	}
	os.Exit(code)
}

func Mkdir(name string, perm os.FileMode) error {
	if VarOsMock == nil {
		return VarOsMock.Mkdir(name, perm)
	}
	return os.Mkdir(name, perm)
}

func Chdir(dir string) error {
	if VarOsMock == nil {
		return VarOsMock.Chdir(dir)
	}
	return os.Chdir(dir)
}

func SameFile(fi1, fi2 os.FileInfo) bool {
	if VarOsMock == nil {
		return VarOsMock.SameFile(fi1, fi2)
	}
	return os.SameFile(fi1, fi2)
}

func Getgroups() ([]int, error) {
	if VarOsMock == nil {
		return VarOsMock.Getgroups()
	}
	return os.Getgroups()
}

func Getpid() int {
	if VarOsMock == nil {
		return VarOsMock.Getpid()
	}
	return os.Getpid()
}

func TempDir() string {
	if VarOsMock == nil {
		return VarOsMock.TempDir()
	}
	return os.TempDir()
}

func Remove(name string) error {
	if VarOsMock == nil {
		return VarOsMock.Remove(name)
	}
	return os.Remove(name)
}

func Geteuid() int {
	if VarOsMock == nil {
		return VarOsMock.Geteuid()
	}
	return os.Geteuid()
}
