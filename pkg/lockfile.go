package pkg

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sawadashota/gorew/driver"
)

type lockfile struct {
	path string
}

func NewLockfile(driver driver.Driver) (*lockfile, error) {
	l := &lockfile{path: driver.Configuration().LockfilePath()}
	if err := l.init(); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *lockfile) List() ([]Package, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pkgs []Package
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		pkgs = append(pkgs, newGoPackage(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return pkgs, nil
}

func (l *lockfile) exist(pkg Package) (bool, error) {
	pkgs, err := l.List()
	if err != nil {
		return false, err
	}

	for _, p := range pkgs {
		if p.Source() == pkg.Source() {
			return true, nil
		}
	}

	return false, nil
}

func (l *lockfile) Add(pkg Package) error {
	exist, err := l.exist(pkg)
	if err != nil {
		return err
	}

	if exist {
		return errors.Errorf("%s has already been added", pkg.Source())
	}

	if err := pkg.Install(); err != nil {
		return err
	}

	file, err := os.OpenFile(l.path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(pkg.Source() + "\n")
	return err
}

func (l *lockfile) Remove(pkg Package) error {
	exist, err := l.exist(pkg)
	if err != nil {
		return err
	}

	if !exist {
		return errors.Errorf("%s is not exists", pkg.Source())
	}

	pkgs, err := l.List()
	if err != nil {
		return err
	}

	var b bytes.Buffer
	for _, p := range pkgs {
		if p.Source() != pkg.Source() {
			b.WriteString(p.Source())
			b.WriteString("\n")
		}
	}

	file, err := os.Create(l.path)
	if err != nil {
		return err
	}

	_, err = file.Write(b.Bytes())
	return err
}

func (l *lockfile) init() error {
	fileInfo, err := os.Stat(l.path)

	if os.IsNotExist(err) {
		return l.create()
	} else if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return errors.Errorf("%s is directory", l.path)
	}

	return nil
}

func (l *lockfile) create() error {
	dir := filepath.Dir(l.path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(l.path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return file.Close()
}
