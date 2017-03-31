package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/daiLlew/go-run-it/model"
	"github.com/daiLlew/go-run-it/util"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
)

func main() {
	targetEnv := flag.String("env", "", "Flag to specify which environment you want to start up.")
	flag.Parse()
	ws := loadWorkspace(*targetEnv)

	sigChan := make(chan  os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		sig := <-sigChan
		switch sig {
		default:
			gracefulShutdown(ws)
		}
	}()

	util.Exec(ws)
}

func loadWorkspace(env string) *model.Workspace {
	envConfig := path.Join("environments", env+".json")

	if _, err := os.Stat(envConfig); os.IsNotExist(err) {
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(envConfig)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	var ws model.Workspace
	if err := json.Unmarshal(bytes, &ws); err != nil {
		os.Exit(0)
	}
	return &ws
}

func gracefulShutdown(ws *model.Workspace) {
	for _, app := range ws.Apps {
		fmt.Printf("Closing log file: %s", app.LogFile.Name())
		app.LogFile.Close()
	}
}