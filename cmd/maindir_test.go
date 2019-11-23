package cmd

import (
	"path"
	"testing"

	"github.com/pkg/profile"
)

func TestMainDirs(t *testing.T) {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("../pprof/maindir"), profile.NoShutdownHook).Stop()

	gitRoot := path.Join(getSrcPathFromEnv(), "github.com/sawadashota/gorew")

	dirs := mainDirs(gitRoot)

	if len(*dirs) != 1 {
		t.Errorf("\"github.com/sawadashota/gorew\" has 1 main.go. Acutual: %v\n", *dirs)
	}

	if (*dirs)[0] != gitRoot {
		t.Errorf("mainDirs should return  %v. Acutual: %v\n", gitRoot, (*dirs)[0])
	}
}
