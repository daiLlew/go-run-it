package main

import (
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/go-ns/log"
	"github.com/daiLlew/go-run-it/application"
	"github.com/daiLlew/go-run-it/model"
	"github.com/daiLlew/go-run-it/util"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	var executor application.TaskExecutor

	if runtime.GOOS != "darwin" {
		fmt.Println("Currently only OS X is supported.")
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile("project-config.json")
	if err != nil {
		log.ErrorC("Failed to load config.", err, nil)
		os.Exit(0)
	}

	var env model.Environment
	if err := json.Unmarshal(bytes, &env); err != nil {
		log.ErrorC("Failed to unmarshal JSON bytes.", err, nil)
		os.Exit(0)
	}

	executor = &application.DarwinTaskExecutor{CmdFactory: &util.DarwinCmdFactory{}}
	executor.Exec(env)
}
