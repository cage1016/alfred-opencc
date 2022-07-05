package main

import (
	"errors"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		wf.FatalError(errors.New("please provide some input 👀"))
	}

	handlers := map[string]func(*aw.Workflow, []string) error{}

	h, found := handlers[args[0]]
	if !found {
		wf.FatalError(errors.New("command not recognized 👀"))
	}

	err := h(wf, args[1:])
	if err != nil {
		wf.FatalError(err)
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
