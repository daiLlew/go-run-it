package handler

import (
	"encoding/json"
	"github.com/daiLlew/go-run-it/model"
	"net/http"
	"strings"
)

type StatusHandler struct {
	Workspace model.Workspace
}

type StatusResponse struct {
	Statuses []*AppStatus
}

type AppStatus struct {
	ID      string
	Name    string
	IsAlive bool
	PID     int
}

func (s *StatusResponse) AddAppStatus(app *AppStatus) {
	s.Statuses = append(s.Statuses, app)
}

func (s *StatusHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	statusResponse := &StatusResponse{
		Statuses: make([]*AppStatus, 0),
	}

	for _, app := range s.Workspace.Apps {
		var pid = 0
		var isAlive = false
		if runProcess, ok := app.Processes[model.RUN_TASK_NAME]; ok {
			isAlive = true
			pid = runProcess.Process.Pid
		}

		statusResponse.AddAppStatus(&AppStatus{
			ID:      strings.Replace(strings.ToLower(app.Name), " ", "-", -1),
			Name:    app.Name,
			IsAlive: isAlive,
			PID:     pid,
		})
	}

	json.NewEncoder(w).Encode(statusResponse)
}
