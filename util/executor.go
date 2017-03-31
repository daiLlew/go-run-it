package util

import (
	"fmt"
	"github.com/daiLlew/go-run-it/model"
	"log"
	"sync"
)

const MAGENTA = "\033[35m"
const RESET_COLOUR = "\033[0m"
const LIGHT_BLUE = "\033[96m"
const LIGHT_GREEN = "\033[92m"
const LIGHT_YELLOW = "\033[93m"
const START_CMD_MSG_FMT = "%s[start-up]%s %sSuccessfully executed '%s'\n%s"

func Exec(ws *model.Workspace) {
	ws.CleanUpLogFiles()

	var wg sync.WaitGroup
	for _, app := range ws.Apps {
		wg.Add(1)

		go func(app *model.Application) {
			defer wg.Done()

			for name, cmd := range app.GenerateCommands(app.CreateLogFile()) {
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
				fmt.Printf(START_CMD_MSG_FMT, LIGHT_GREEN, RESET_COLOUR, LIGHT_BLUE, name, RESET_COLOUR)
			}

		}(app)
	}
	wg.Wait()
	fmt.Println("Alls apps have now started.")
}
