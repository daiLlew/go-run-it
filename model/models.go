package model

import (
	"fmt"
	"os/exec"
)

type Task struct {
	CMDS []*exec.Cmd
}

func (t *Task) Run() {
	for _, c := range t.CMDS {
		fmt.Printf("%s\n", c.Args)
		c.Run()
	}
}

type Environment struct {
	Applications []Application
	cmds         []*exec.Cmd
}

type Application struct {
	Name      string `json:"name,omitempty"`
	Directory string `json:"directory,omitempty"`
	Commands  []string `json:"commands,omitempty"`
	URL       string `json:"url,omitempty"`
}
