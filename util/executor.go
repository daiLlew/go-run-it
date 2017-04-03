package util

import (
	"fmt"
	"github.com/daiLlew/go-run-it/logger"
	"github.com/daiLlew/go-run-it/model"
	"log"
	"sync"
)

func Exec(ws *model.Workspace) {
	ws.CleanUpLogFiles()

	var wg sync.WaitGroup
	for _, app := range ws.Apps {
		wg.Add(1)

		go func(app *model.Application) {
			defer wg.Done()

			buildCMD, runCMD := app.GenerateCommands(app.CreateLogFile())

			if buildCMD != nil {
				if err := buildCMD.Run(); err != nil {
					log.Fatal(err)
				}
				logger.StartupDebug("Successfully executed build.")
			}

			if err := runCMD.Start(); err != nil {
				log.Fatal(err)
			}
			logger.StartupDebug(fmt.Sprintf("Successfully executed %s run cmd PID %d", app.Name, runCMD.Process.Pid))
		}(app)
	}
	wg.Wait()
	logger.StartupDebug("All start-up tasks have successfully executed.")
}
