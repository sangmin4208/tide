package main

import (
	"fmt"

	"github.com/sangmin4208/tidy/config"
	"github.com/sangmin4208/tidy/lib"
	"github.com/sangmin4208/tidy/model"
)

func main() {
	date := lib.Prompt()
	codes := lib.ReadJsonFile(config.CODE_FILE)
	lib.MakeDirectory(config.OUTPUT_PATH)
	files := []model.File{}
	for _, code := range codes {
		for i := date.StartMonth; i <= date.EndMonth; i++ {
			fname := ""
			if i < 10 {
				fname = fmt.Sprintf("%v/0%v", date.Year, i)
			} else {
				fname = fmt.Sprintf("%v/%v", date.Year, i)
			}
			lib.MakeDirectory(fmt.Sprintf("%s/%s", config.OUTPUT_PATH, fname))
			URI := lib.GetURI(config.DOWNLOAD_URI, code, date.Year, i)
			data := lib.GetData(URI)
			data = lib.DeleteLine(data, 5)
			files = append(files, model.File{Path: fmt.Sprintf("%s/%s.txt", fname, code.Name), Body: data})
		}
	}

	for _, file := range files {
		lib.WriteFile(file)
	}
}
