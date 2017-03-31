package application
/*
import (
	"github.com/ONSdigital/go-ns/log"
	"github.com/daiLlew/go-run-it/model"
	"github.com/daiLlew/go-run-it/util"
	"os"
	"sync"
)

type TaskExecutor interface {
	Exec(env model.Environment)
}

type DarwinTaskExecutor struct {
	CmdFactory util.CmdFactory
}

func (d DarwinTaskExecutor) Exec(env model.Environment) {

	d.CmdFactory.OpenTerminalTabsCommand(len(env.Applications)).Run()

	appTasks := make([]*model.Task, 0)

	for i, app := range env.Applications {
		appTasks = append(appTasks, d.CmdFactory.GenerateAppTasks(i+1, app))
	}

	var wg sync.WaitGroup
	wg.Add(len(appTasks))

	for _, tasks := range appTasks {
		go runApplication(tasks, &wg)
	}
	wg.Wait()
}

func runApplication(task *model.Task, wg *sync.WaitGroup) {
	for _, cmd := range task.CMDS {
		log.Debug("Executing command", log.Data{
			"command": cmd.Args,
		})
		err := cmd.Run()
		if err != nil {
			log.ErrorC(err.Error(), err, nil)
			os.Exit(0)
		}
	}
	wg.Done()
}
*/
