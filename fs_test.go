package staticmodtimefs_test

import (
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/lsymds/staticmodtimefs"
)

var tm = time.Now()

func TestReadDirFSImplementationReturnedIfSupported(t *testing.T) {
	f := staticmodtimefs.NewStaticModTimeFS(os.DirFS("./"), tm)

	if _, ok := f.(fs.ReadDirFS); !ok {
		t.Error("fs is not [fs.ReadDirFS]")
	}
}

func TestStaticModTimeReturnedForFiles(t *testing.T) {
	f := staticmodtimefs.NewStaticModTimeFS(os.DirFS("./"), tm)

	if fi, err := f.Open("README.md"); err != nil {
		t.Errorf("could not open file: %e", err)
	} else if s, err := fi.Stat(); err != nil {
		t.Errorf("could not stat file: %e", err)
	} else if s.ModTime() != tm {
		t.Errorf("file ModTime was not %v, got %v", tm, s.ModTime())
	}
}

func TestStaticModTimeReturnedForDirectories(t *testing.T) {
	f := staticmodtimefs.NewStaticModTimeFS(os.DirFS("./"), tm).(fs.ReadDirFS)

	if fi, err := f.ReadDir(".github"); err != nil {
		t.Errorf("could not read dir: %v", err)
	} else if s, err := fi[0].Info(); err != nil {
		t.Errorf("could not get file info: %v", err)
	} else if s.ModTime() != tm {
		t.Errorf("dir ModTime was %v, wanted %v", s.ModTime(), tm)
	}
}
