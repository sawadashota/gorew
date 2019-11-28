package pkg

import (
	"os"
	"os/exec"
)

type Package interface {
	Install() error
	Source() string
}

type GoPackage struct {
	src string
}

func NewGoPackage(src string) *GoPackage {
	return &GoPackage{src: src}
}

func (gp *GoPackage) Install() error {
	cmd := exec.Command("go", "get", gp.src)
	cmd.Env = append(os.Environ(), "GO111MODULE=off")
	return cmd.Run()
}

func (gp *GoPackage) Source() string {
	return gp.src
}
