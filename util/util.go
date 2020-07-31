package util

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/saanai/util-sys/data"
	"github.com/saanai/util-sys/entity"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func MapEntitiesToDataThreads(entities []entity.Thread) (data []data.Thread) {
	for _, entity := range entities {
		data = append(data, *mapEntityToDataThread(&entity))
	}
	return data
}

func mapEntityToDataThread(entity *entity.Thread) (data *data.Thread) {
	return data.NewDataThread(entity.Id, entity.Uuid, entity.Topic, entity.UserId, entity.CreatedAt)
}
