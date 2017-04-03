package handler

import (
	"fmt"
	"github.com/daiLlew/go-run-it/logger"
	"github.com/daiLlew/go-run-it/model"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type ApplicationHandler struct {
	workspace *model.Workspace
}

// /stop/{name}/{PID}
func StopHandlerFunc(ws *model.Workspace) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pid := r.URL.Query().Get(":PID")
		targetPID, _ := strconv.Atoi(pid)
		appID := r.URL.Query().Get(":appID")
		var targetApp *model.Application = nil

		for _, a := range ws.Apps {
			if a.ID == appID {
				targetApp = a
				break
			}
		}
		if targetApp == nil {
			writeErrorResponse(w, fmt.Sprintf("Failed to stop application '%s' as it does not exist.", appID), 500)
			return
		}

		for k, v := range targetApp.Processes {
			if v.Process.Pid == targetPID {

				logger.StartupDebug(fmt.Sprintf("%+v", v.ProcessState))

				args := []string{"-o", "ppid= ", strconv.Itoa(v.Process.Pid)}
				bytes, _ := exec.Command("ps", args...).Output()

				ppid, _ := strconv.Atoi(strings.TrimSpace(string(bytes)))
				fmt.Printf("Parent ID=%d\n", ppid)

				if err := syscall.Kill(-v.Process.Pid, syscall.SIGKILL); err != nil {
					writeErrorResponse(w, fmt.Sprintf("Failed to stop application process '%d' as it does not exist.", targetPID), 500)
					return
				}
				delete(targetApp.Processes, k)
				w.Write([]byte(fmt.Sprintf("Successfully to stopped application '%s' with PID '%d'.", appID, targetPID)))
				w.WriteHeader(200)
				return
			}
		}
	}
}

func writeErrorResponse(w http.ResponseWriter, msg string, status int) {
	logger.LogError(msg)
	w.Write([]byte(msg))
	w.WriteHeader(status)
}
