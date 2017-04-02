package webModel

import (
	"github.com/daiLlew/go-run-it/model"
	"strings"
)

type WorkSpace struct {
	Name string
	Apps []*App
}

type App struct {
	ID   string
	Name string
	URL  string
}

func Convert(domainModel *model.Workspace) *WorkSpace {

	result := &WorkSpace{
		Apps: make([]*App, 0),
		Name: domainModel.Name,
	}

	for _, x := range domainModel.Apps {
		result.Apps = append(result.Apps, &App{
			ID: strings.Replace(strings.ToLower(x.Name), " ", "-", -1),
			Name: x.Name,
			URL:  x.URL,
		})
	}
	return result
}
