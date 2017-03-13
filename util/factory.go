package util

import (
	"github.com/daiLlew/go-run-it/model"
)

type CmdFactory interface {
	OpenTerminalTabsCommand(number int) *model.Task

	GenerateAppTasks(tabIndex int, app model.Application) *model.Task
}