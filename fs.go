package staticmodtimefs

import (
	"errors"
	"io/fs"
	"time"
)

// NewStaticModTimeFS creates a wrapper filesystem around an existing [fs.FS] or [fs.ReadDirFS] implementation where
// any modification times returned are always of the provided constant value.
func NewStaticModTimeFS(f fs.FS, modified time.Time) fs.FS {
	sfs := &staticModTimeFS{fs: f, modified: &modified}

	if _, ok := f.(fs.ReadDirFS); ok {
		return &staticModTimeReadDirFS{staticModTimeFS: sfs}
	} else {
		return sfs
	}
}

type staticModTimeFS struct {
	fs       fs.FS
	modified *time.Time
}

func (s *staticModTimeFS) Open(name string) (fs.File, error) {
	f, err := s.fs.Open(name)
	if err != nil {
		return nil, err
	}

	return &staticModTimeFile{File: f, modified: s.modified}, nil
}

type staticModTimeFile struct {
	fs.File
	modified *time.Time
}

func (s *staticModTimeFile) Stat() (fs.FileInfo, error) {
	i, err := s.File.Stat()
	if err != nil {
		return nil, err
	}

	return &staticModTimeFileInfo{FileInfo: i, modified: s.modified}, nil
}

type staticModTimeFileInfo struct {
	fs.FileInfo
	modified *time.Time
}

func (s *staticModTimeFileInfo) ModTime() time.Time {
	return *s.modified
}

type staticModTimeReadDirFS struct {
	*staticModTimeFS
}

func (s *staticModTimeReadDirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	f, ok := s.fs.(fs.ReadDirFS)
	if !ok {
		return nil, errors.New("not implemented")
	}

	dirs, err := f.ReadDir(name)
	if err != nil {
		return nil, err
	}

	es := make([]fs.DirEntry, len(dirs))
	for i, d := range dirs {
		es[i] = &staticModTimeDirEntry{DirEntry: d, modified: s.modified}
	}

	return es, nil
}

type staticModTimeDirEntry struct {
	fs.DirEntry
	modified *time.Time
}

func (d *staticModTimeDirEntry) Info() (fs.FileInfo, error) {
	fi, err := d.DirEntry.Info()
	if err != nil {
		return nil, err
	}

	return &staticModTimeFileInfo{FileInfo: fi, modified: d.modified}, nil
}
