package util

import (
	"fmt"
	"github.com/daiLlew/go-run-it/model"
	"os/exec"
)

const OSA_BASE = "tell application \"Terminal\" to "
const OSA_DO_SCRIPT = OSA_BASE + "do script \"%s\" in tab %d of front window"
const OPEN_NEW_TERMINAL_TAB = "tell application \"System Events\" to keystroke \"t\" using {command down}"
const RENAME_TAB = OSA_BASE + "set custom title of tab %d of window 1 to \"%s\""


type DarwinCmdFactory struct {
}

func (f *DarwinCmdFactory) OpenTerminalTabsCommand(number int) *model.Task {
	var commands = make([]*exec.Cmd, 0)
	commands = append(commands, makeCommand(OSA_BASE + "activate"))

	for i := 0; i < number; i++ {
		commands = append(commands, makeCommand(OPEN_NEW_TERMINAL_TAB))
	}
	return &model.Task{CMDS: commands}
}

func (f *DarwinCmdFactory) GenerateAppTasks(tabIndex int, app model.Application) *model.Task {
	var commands = make([]*exec.Cmd, 0)
	commands = append(commands, makeCommand(fmt.Sprintf(OSA_DO_SCRIPT, "cd "+app.Directory, tabIndex)))
	commands = append(commands, makeCommand(fmt.Sprintf(RENAME_TAB, tabIndex, app.Name)))

	for _, c := range app.Commands {
		commands = append(commands, makeCommand(fmt.Sprintf(OSA_DO_SCRIPT, c, tabIndex)))
	}

	return &model.Task{commands}
}

func makeCommand(commandStr string) *exec.Cmd {
	return exec.Command("osascript", "-e", commandStr)
}
