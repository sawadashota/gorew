package cmd

import (
	"path"
	"testing"
)

func TestMainDirs(t *testing.T) {
	gitRoot := path.Join(getSrcPathFromEnv(), "github.com/sawadashota/gocmd")

	dirs := mainDirs(gitRoot)

	if len(dirs) != 1 {
		t.Errorf("\"github.com/sawadashota/gocmd\" has 1 main.go. Acutual: %v\n", dirs)
	}

	if dirs[0] != gitRoot {
		t.Errorf("mainDirs shouldf return  %v. Acutual: %v\n", gitRoot, dirs[0])
	}
}
