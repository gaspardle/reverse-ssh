package main

import (
	"github.com/judwhite/go-svc"
)

func (prg *program) Init(env svc.Environment) error {
	return nil
}

func (prg *program) Start() error {
	go run(prg.p, prg.server)
	return nil
}

func (prg *program) Stop() error {
	prg.server.Close()
	return nil
}
