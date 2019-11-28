package pkg

import "os/exec"

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
	return exec.Command("GO111MODULE=off", "go", "get", gp.src).Run()
}

func (gp *GoPackage) Source() string {
	return gp.src
}
