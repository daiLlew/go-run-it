package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/daiLlew/go-run-it/handler"
	"github.com/daiLlew/go-run-it/model"
	"github.com/daiLlew/go-run-it/util"
	"github.com/daiLlew/go-run-it/webModel"
	"github.com/gorilla/pat"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
)

func main() {
	targetEnv := flag.String("env", "ws", "Flag to specify which environment you want to start up.")
	flag.Parse()
	ws := loadWorkspace(*targetEnv)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		sig := <-sigChan
		switch sig {
		default:
			gracefulShutdown(ws)
		}
	}()

	util.Exec(ws)

	router := pat.New()

	homepageHandler := &handler.TemplateHandler{
		Filename:  "homepage.html",
		Workspace: webModel.Convert(ws),
	}

	statusHandler := &handler.StatusHandler{Workspace: *ws}

	router.Get("/status", statusHandler.GetStatus)
	router.Get("/", homepageHandler.ServeHTTP)

	if err := http.ListenAndServe(":9001", router); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
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
	ws.Shutdown()
	os.Exit(0)
}