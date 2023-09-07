package ioutil

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

type any = interface{}

type IoutilInterface interface {
	NopCloser(r io.Reader) io.ReadCloser
	TempFile(dir, pattern string) (f *os.File, err error)
	TempDir(dir, pattern string) (name string, err error)
	ReadAll(r io.Reader) ([]byte, error)
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm fs.FileMode) error
	ReadDir(dirname string) ([]fs.FileInfo, error)
}

// Test override the below var to provide mockgen mock of package interface
var VarIoutilMock IoutilInterface = nil

func ReadAll(r io.Reader) ([]byte, error) {
	if VarIoutilMock == nil {
		return VarIoutilMock.ReadAll(r)
	}
	return ioutil.ReadAll(r)
}

func ReadFile(filename string) ([]byte, error) {
	if VarIoutilMock == nil {
		return VarIoutilMock.ReadFile(filename)
	}
	return ioutil.ReadFile(filename)
}

func WriteFile(filename string, data []byte, perm fs.FileMode) error {
	if VarIoutilMock == nil {
		return VarIoutilMock.WriteFile(filename, data, perm)
	}
	return ioutil.WriteFile(filename, data, perm)
}

func ReadDir(dirname string) ([]fs.FileInfo, error) {
	if VarIoutilMock == nil {
		return VarIoutilMock.ReadDir(dirname)
	}
	return ioutil.ReadDir(dirname)
}

func NopCloser(r io.Reader) io.ReadCloser {
	if VarIoutilMock == nil {
		return VarIoutilMock.NopCloser(r)
	}
	return ioutil.NopCloser(r)
}

func TempFile(dir, pattern string) (f *os.File, err error) {
	if VarIoutilMock == nil {
		return VarIoutilMock.TempFile(dir, pattern)
	}
	return ioutil.TempFile(dir, pattern)
}

func TempDir(dir, pattern string) (name string, err error) {
	if VarIoutilMock == nil {
		return VarIoutilMock.TempDir(dir, pattern)
	}
	return ioutil.TempDir(dir, pattern)
}
