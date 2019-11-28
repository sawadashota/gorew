package pkg_test

import (
	"os"
	"testing"

	"github.com/sawadashota/gorew/driver/configuration"
	"github.com/sawadashota/gorew/pkg"
)

type testConfiguration struct {
	lockfilePath string
}

func newTestConfiguration(lockfilePath string) *testConfiguration {
	return &testConfiguration{lockfilePath: lockfilePath}
}

func (t *testConfiguration) LockfilePath() string {
	return t.lockfilePath
}

type testDriver struct {
	c configuration.Provider
}

func newTestDriver(lockfilePath string) *testDriver {
	return &testDriver{c: newTestConfiguration(lockfilePath)}
}

func (d *testDriver) Configuration() configuration.Provider {
	return d.c
}

func TestLockfile(t *testing.T) {
	if _, err := os.Stat("./testdata"); os.IsExist(err) {
		if err := os.RemoveAll("./testdata"); err != nil {
			t.Fatal(err)
		}
	}

	d := newTestDriver("./testdata/.gorew")
	lockfile, err := pkg.NewLockfile(d)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat("./testdata/.gorew"); os.IsNotExist(err) {
		t.Error("expected ./testdata/.gorew exists but un-exists")
	}

	if err := lockfile.Add(pkg.NewGoPackage("github.com/sawadashota/gorew")); err != nil {
		t.Error(err)
	}

	if err := lockfile.Add(pkg.NewGoPackage("github.com/sawadashota/gorew")); err == nil {
		t.Error("same package should be rejected")
	}

	pkgs, err := lockfile.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs) != 1 {
		t.Errorf("expected 1 packages but %d", len(pkgs))
	}

	if err := lockfile.Remove(pkg.NewGoPackage("github.com/sawadashota/gorew")); err != nil {
		t.Error(err)
	}

	if err := lockfile.Remove(pkg.NewGoPackage("github.com/sawadashota/gorew")); err == nil {
		t.Error("un-exists package should be rejected")
	}
}
