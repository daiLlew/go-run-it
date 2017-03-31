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

const MAGENTA = "\033[35m"
const NO_COL = "\033[0m"
const BLUE = "\033[96m"

const GREEN = "\033[92m"
const YELLOW = "\033[93m"

const SHUT_DOWN_MSG_FMT = "%s[shutdown]%s %s%s%s\n"

type Workspace struct {
	Name string `json:"name,omitempty"`
	Apps []*Application `json:"apps,omitempty"`
}

type Application struct {
	Name      string `json:"name,omitempty"`
	URL       string `json:"url,omitempty"`
	Dir       string `json:"dir,omitempty"`
	Tasks     []Command `json:"tasks,omitempty"`
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
		fmt.Printf(formatShutdownMessage(fmt.Sprintf("Closed file: %s.", app.LogFile.Name())))
		app.LogFile.Close()
		for n, p := range app.Processes {
			fmt.Printf(formatShutdownMessage(fmt.Sprintf("Terminted Process %d '%s'.", p.Process.Pid, n)))
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

func (a *Application) GenerateCommands(outs ...io.Writer) map[string]*exec.Cmd {
	a.Processes = make(map[string]*exec.Cmd, 0)

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
		a.Processes[task.Name] = cmd
	}
	return a.Processes
}

func formatShutdownMessage(msg string) string {
	return fmt.Sprintf(SHUT_DOWN_MSG_FMT, YELLOW, NO_COL, BLUE, msg, NO_COL)
}
