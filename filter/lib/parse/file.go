package parse

import (
	"github.com/sangmin4208/filter/model"
)

func File(m map[string][]model.Tide) []*model.File {
	files := []*model.File{}
	for key := range m {
		file := &model.File{}
		file.Name = key
		files = append(files, file)
	}

	for _, file := range files {
		body := ""
		for _, tide := range m[file.Name] {
			body += TidetoA(tide) + "\n"
		}
		file.Body = body
	}
	return files
}
