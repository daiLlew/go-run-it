package model

import (
	"fmt"
	"github.com/daiLlew/go-run-it/logger"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const BUILD_TASK_NAME = "build"
const RUN_TASK_NAME = "run"

type Workspace struct {
	Name string `json:"name,omitempty"`
	Apps []*Application `json:"apps,omitempty"`
}

type Application struct {
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	Dir      string `json:"dir,omitempty"`
	BuildCMD *Command `json:"buildCmd,omitempty"`
	RunCMD   *Command `json:"runCmd,omitempty"`

	Processes map[string]*exec.Cmd
	LogFile   *os.File
}

type Command struct {
	Name string `json:"name,omitempty"`
	Dir  string `json:"dir,omitempty"`
	Path string `json:"path,omitempty"`
	Args []string `json:"args,omitempty"`
}

func (ws *Workspace) CleanUpLogFiles() {
	if err := os.RemoveAll("logs"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("logs", 0777); err != nil {
		log.Fatal(err)
	}
}

func (ws *Workspace) Shutdown() {
	fmt.Println("")
	for _, app := range ws.Apps {
		logger.ShutdownDebug(fmt.Sprintf("Closed file: %s.", app.LogFile.Name()))
		app.LogFile.Close()
		for n, p := range app.Processes {
			logger.ShutdownDebug(fmt.Sprintf("Terminted Process %d '%s'.", p.Process.Pid, n))
			p.Process.Kill()
		}
	}
	fmt.Println("")
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
		log.Fatal(err.Error())
	}
	a.LogFile = f
	return f
}

func (a *Application) GenerateCommands(outs ...io.Writer) (build *exec.Cmd, run *exec.Cmd) {
	a.Processes = make(map[string]*exec.Cmd, 0)
	var buildCMD *exec.Cmd = nil

	if a.BuildCMD != nil {
		buildCMD = createCommand(a.BuildCMD, a, BUILD_TASK_NAME, outs...)
	}
	return buildCMD, createCommand(a.RunCMD, a, RUN_TASK_NAME, outs...)
}

func createCommand(c *Command, a *Application, taskName string, outs ...io.Writer) *exec.Cmd {
	cmd := exec.Command(c.Path, c.Args...)
	mw := io.MultiWriter(outs...)
	cmd.Stdout = mw
	cmd.Stderr = mw

	if len(c.Dir) > 0 {
		cmd.Dir = c.Dir
	} else {
		cmd.Dir = a.Dir
	}
	a.Processes[taskName] = cmd
	return cmd
}
