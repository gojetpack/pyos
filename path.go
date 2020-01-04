package pyos

import (
	"errors"
	"os"
)

type FileType int

// See: https://en.wikipedia.org/wiki/Unix_file_types
const (
	FileTypeUnknown     = FileType(-2)
	FileTypeNotExist    = FileType(-1)
	FileTypeRegularFile = FileType(0)
	FileTypeDirectory   = FileType(1)
)

type path struct {
}

// Return True if path refers to an existing path or an open file descriptor.
// Returns False for broken symbolic links.
// On some platforms, this function may return False if permission is not granted
// to execute os.stat() on the requested file, even if the path physically exists.
// see: https://docs.python.org/3/library/os.path.html#os.path.exists
func (p path) Exist(path string) bool {
	_, exist, _ := p.PathInfo(path)
	return exist
}

func (path) PathInfo(path string) (os.FileInfo, bool, error) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fi, false, err
	}
	return fi, err == nil, err
}

func (p path) GetFileType(path string) (FileType, error) {
	fi, exist, err := p.PathInfo(path)
	if !exist {
		return FileTypeNotExist, err
	}
	if err != nil {
		return FileTypeUnknown, err
	}
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		// do file stuff
		return FileTypeRegularFile, nil
	case mode.IsDir():
		// do directory stuff
		return FileTypeDirectory, nil
	}
	return FileTypeUnknown, errors.New("unknown file type")
}

// Return True if path is an existing regular file.
// This follows symbolic links, so both islink() and isfile() can be true for the same path.
// see: https://docs.python.org/3/library/os.path.html#os.path.isdir
func (p path) IsFile(path string) bool {
	fileType, err := p.GetFileType(path)
	if err != nil {
		return false
	}
	return fileType == FileTypeRegularFile
}

// Return True if path is an existing directory.
// This follows symbolic links, so both islink() and isdir() can be true for the same path.
// see:
func (p path) IsDir(path string) bool {
	fileType, err := p.GetFileType(path)
	if err != nil {
		return false
	}
	return fileType == FileTypeDirectory
}
