package model

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Workspace struct {
	Name string `json:"name,omitempty"`
	Apps []*Application `json:"apps,omitempty"`
}

type Application struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	Dir     string `json:"dir,omitempty"`
	Tasks   []Command `json:"tasks,omitempty"`
	LogFile *os.File
}

type Command struct {
	Dir  string `json:"dir,omitempty"`
	Path string `json:"path,omitempty"`
	Args []string `json:"args,omitempty"`
}

func (ws *Workspace) CleanUpLogFiles() {
	if err := os.RemoveAll("logs"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted logs dir")
	if err := os.Mkdir("logs", 0777); err != nil {
		log.Fatal(err)
	}
	fmt.Println("created logs dir")
}

func (a *Application) CreateLogFile() *os.File {
	filename := path.Join("logs", strings.ToLower(strings.Replace(a.Name, " ", "-", -1))) + ".log"
	_, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	workingDir, _ := os.Getwd()
	logFilePath := path.Join(workingDir, filename)

	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("OpenLogFile fail")
		log.Fatal(err.Error())
	}
	a.LogFile = f
	return f
}

func (a *Application) GenerateCommands(outs...io.Writer) []*exec.Cmd {
	cmds := make([]*exec.Cmd, 0)

	for _, task := range a.Tasks {
		cmd := exec.Command(task.Path, task.Args...)
		mw := io.MultiWriter(outs...)
		cmd.Stdout = mw
		cmd.Stderr = mw

		if len(task.Dir) > 0 {
			cmd.Dir = task.Dir
		} else {
			cmd.Dir = a.Dir
		}

		cmds = append(cmds, cmd)
	}
	return cmds
}