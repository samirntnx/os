# Go OS Wrapper for Unit Testing

This Go package provides a wrapper for the standard `os` package to simplify unit testing when working with file and directory operations. It allows you to mock file system operations, making it easier to write test cases for code that interacts with the file system.

## Installation

To use this package in your Go project, you can simply import it using:

```go
import (
    "github.com/samirntnx/os"
    "github.com/samirntnx/os/exec"
)

func main() {
    os.FS.Openfile("testfile")
    exec.EXEC.Command("ls")
}
```

## For existing codes
User's gopatch tool can be used to convert existing code base to use this repo.

os.patch
```
@@
@@
-import "os"
+import "os"
+import gos "github.com/samirntnx/os"

os
@@
@@
-os.ExpandEnv(...)
+gos.Fn.ExpandEnv(...)

@@
@@
-os.Symlink(...)
+gos.Fn.Symlink(...)

@@
@@
-os.Getuid(...)
+gos.Fn.Getuid(...)

@@
@@
-os.Executable(...)
+gos.Fn.Executable(...)

@@
@@
-os.Chown(...)
+gos.Fn.Chown(...)

@@
@@
-os.WriteFile(...)
+gos.Fn.WriteFile(...)

@@
@@
-os.NewFile(...)
+gos.Fn.NewFile(...)

@@
@@
-os.MkdirAll(...)
+gos.Fn.MkdirAll(...)

@@
@@
-os.IsPermission(...)
+gos.Fn.IsPermission(...)

@@
@@
-os.UserHomeDir(...)
+gos.Fn.UserHomeDir(...)

@@
@@
-os.DirFS(...)
+gos.Fn.DirFS(...)

@@
@@
-os.Lstat(...)
+gos.Fn.Lstat(...)

@@
@@
-os.IsNotExist(...)
+gos.Fn.IsNotExist(...)

@@
@@
-os.TempDir(...)
+gos.Fn.TempDir(...)

@@
@@
-os.Chdir(...)
+gos.Fn.Chdir(...)

@@
@@
-os.Open(...)
+gos.Fn.Open(...)

@@
@@
-os.Chmod(...)
+gos.Fn.Chmod(...)

@@
@@
-os.Truncate(...)
+gos.Fn.Truncate(...)

@@
@@
-os.CreateTemp(...)
+gos.Fn.CreateTemp(...)

@@
@@
-os.Setenv(...)
+gos.Fn.Setenv(...)

@@
@@
-os.Getpid(...)
+gos.Fn.Getpid(...)

@@
@@
-os.RemoveAll(...)
+gos.Fn.RemoveAll(...)

@@
@@
-os.Hostname(...)
+gos.Fn.Hostname(...)

@@
@@
-os.Getpagesize(...)
+gos.Fn.Getpagesize(...)

@@
@@
-os.ReadDir(...)
+gos.Fn.ReadDir(...)

@@
@@
-os.Getwd(...)
+gos.Fn.Getwd(...)

@@
@@
-os.NewSyscallError(...)
+gos.Fn.NewSyscallError(...)

@@
@@
-os.Mkdir(...)
+gos.Fn.Mkdir(...)

@@
@@
-os.Remove(...)
+gos.Fn.Remove(...)

@@
@@
-os.Link(...)
+gos.Fn.Link(...)

@@
@@
-os.Exit(...)
+gos.Fn.Exit(...)

@@
@@
-os.Clearenv(...)
+gos.Fn.Clearenv(...)

@@
@@
-os.IsTimeout(...)
+gos.Fn.IsTimeout(...)

@@
@@
-os.Pipe(...)
+gos.Fn.Pipe(...)

@@
@@
-os.MkdirTemp(...)
+gos.Fn.MkdirTemp(...)

@@
@@
-os.SameFile(...)
+gos.Fn.SameFile(...)

@@
@@
-os.LookupEnv(...)
+gos.Fn.LookupEnv(...)

@@
@@
-os.StartProcess(...)
+gos.Fn.StartProcess(...)

@@
@@
-os.FindProcess(...)
+gos.Fn.FindProcess(...)

@@
@@
-os.OpenFile(...)
+gos.Fn.OpenFile(...)

@@
@@
-os.Rename(...)
+gos.Fn.Rename(...)

@@
@@
-os.Getgid(...)
+gos.Fn.Getgid(...)

@@
@@
-os.Getenv(...)
+gos.Fn.Getenv(...)

@@
@@
-os.Getppid(...)
+gos.Fn.Getppid(...)

@@
@@
-os.Expand(...)
+gos.Fn.Expand(...)

@@
@@
-os.Getgroups(...)
+gos.Fn.Getgroups(...)

@@
@@
-os.Create(...)
+gos.Fn.Create(...)

@@
@@
-os.Getegid(...)
+gos.Fn.Getegid(...)

@@
@@
-os.UserCacheDir(...)
+gos.Fn.UserCacheDir(...)

@@
@@
-os.Geteuid(...)
+gos.Fn.Geteuid(...)

@@
@@
-os.IsExist(...)
+gos.Fn.IsExist(...)

@@
@@
-os.Chtimes(...)
+gos.Fn.Chtimes(...)

@@
@@
-os.Readlink(...)
+gos.Fn.Readlink(...)

@@
@@
-os.IsPathSeparator(...)
+gos.Fn.IsPathSeparator(...)

@@
@@
-os.Unsetenv(...)
+gos.Fn.Unsetenv(...)

@@
@@
-os.Environ(...)
+gos.Fn.Environ(...)

@@
@@
-os.Lchown(...)
+gos.Fn.Lchown(...)

@@
@@
-os.Stat(...)
+gos.Fn.Stat(...)

@@
@@
-os.UserConfigDir(...)
+gos.Fn.UserConfigDir(...)

@@
@@
-os.ReadFile(...)
+gos.Fn.ReadFile(...)

@@
@@
-os.LinkError
+gos.LinkError

@@
@@
-os.File
+gos.File

@@
@@
-os.SyscallError
+gos.SyscallError

@@
@@
-os.Process
+gos.Process

@@
@@
-os.ProcessState
+gos.ProcessState


```