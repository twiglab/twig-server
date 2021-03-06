package fcgi

import (
	"context"
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/twiglab/twig"
)

// FcgiServant fcgi实现
type FcgiServant struct {
	file string
	ln   net.Listener
	twig *twig.Twig
}

func NewFcgiServant(file string) *FcgiServant {
	return &FcgiServant{
		file: file,
	}
}

func (s *FcgiServant) Start() (err error) {
	if s.ln, err = net.Listen("unix", s.file); err != nil {
		return
	}

	go func() {
		if err = fcgi.Serve(s.ln, s.twig); err != nil {
			s.twig.Logger.Panic(err)
		}
	}()

	return
}

func (s *FcgiServant) Shutdown(c context.Context) error {
	return s.ln.Close()
}

func (s *FcgiServant) Attach(t *twig.Twig) {
	s.twig = t
}

func (s *FcgiServant) Handler(http.Handler) {
}
