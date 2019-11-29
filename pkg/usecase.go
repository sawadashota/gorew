package pkg

import (
	"fmt"
	"io"
	"os"

	"github.com/sawadashota/gorew/driver"
)

type Handler struct {
	d    driver.Driver
	w    io.Writer
	repo Repository
}

type Option func(*Handler)

func WithDriver(driver driver.Driver) Option {
	return func(handler *Handler) {
		handler.d = driver
	}
}

func WithWriter(writer io.Writer) Option {
	return func(handler *Handler) {
		handler.w = writer
	}
}

func NewHandler(opts ...Option) (*Handler, error) {
	h := &Handler{
		d: driver.NewDefaultDriver(),
		w: os.Stdout,
	}

	for _, opt := range opts {
		opt(h)
	}

	l, err := NewLockfile(h.d)
	if err != nil {
		return nil, err
	}
	h.repo = l

	return h, nil
}

func (h *Handler) List() ([]Package, error) {
	return h.repo.List()
}

func (h *Handler) Add(p Package) error {
	if err := p.Install(); err != nil {
		return err
	}

	return h.repo.Add(p)
}

func (h *Handler) UpdateAll() error {
	ps, err := h.repo.List()
	if err != nil {
		return err
	}

	for _, p := range ps {
		_, _ = fmt.Fprintln(h.w, p.Source())
		err = p.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) Update(p Package) error {
	return p.Update()
}

func (h *Handler) Remove(p Package) error {
	return h.repo.Remove(p)
}

func (h *Handler) InstallAll() error {
	l, err := NewLockfile(h.d)
	if err != nil {
		return err
	}

	ps, err := l.List()
	if err != nil {
		return err
	}

	for _, p := range ps {
		_, _ = h.w.Write([]byte(fmt.Sprintf("Installing %s\n", p.Source())))
		if err := p.Install(); err != nil {
			return err
		}
	}

	return nil
}
