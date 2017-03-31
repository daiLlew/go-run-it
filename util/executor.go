package util

import (
	"fmt"
	"github.com/daiLlew/go-run-it/model"
	"log"
	"os"
	"sync"
)

func Exec(ws *model.Workspace) {
	ws.CleanUpLogFiles()

	var wg sync.WaitGroup
	for _, app := range ws.Apps {
		wg.Add(1)

		go func(app *model.Application) {
			defer wg.Done()

			for _, cmd := range app.GenerateCommands(app.CreateLogFile(), os.Stdout) {
				fmt.Println(app.LogFile)
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
			}
		}(app)
	}
	wg.Wait()
}
