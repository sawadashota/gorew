package cmd

import (
	"path"
	"testing"

	"github.com/pkg/profile"
)

func Test_repo(t *testing.T) {
	srcPath := getSrcPathFromEnv()

	cmd := &commandDir{path.Join(srcPath, "github.com/sawadashota/gorew")}
	if cmd.repo() != "github.com/sawadashota/gorew" {
		t.Errorf("repo should be \"github.com/sawadashota/gorew\". Actual: %v\n", cmd.repo())
	}

}

func Test_basename(t *testing.T) {
	srcPath := getSrcPathFromEnv()

	repoPath := path.Join(srcPath, "github.com/sawadashota/gorew")

	if basename(repoPath) != "gorew" {
		t.Errorf("Basename should be \"gorew\". Actual: %v\n", basename(repoPath))
	}
}

func Test_goCommand(t *testing.T) {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("../pprof/gocomands"), profile.NoShutdownHook).Stop()

	srcPath := getSrcPathFromEnv()
	binPath := getBinPathFromEnv()

	commands := goCommands(srcPath, binPath)

	if len(*commands) < 1 {
		t.Error("go commands should be more than 1")
	}

	if !hasRepo(*commands) {
		t.Error("goCommand() should have \"github.com/sawadashota/gorew\"")
	}
}

func hasRepo(commands []commandDir) bool {
	for _, command := range commands {
		if command.repo() == "github.com/sawadashota/gorew" {
			return true
		}
	}
	return false
}
