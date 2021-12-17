package write

import (
	"os"

	"github.com/sangmin4208/filter/model"
)

func File(path string, file model.File) {
	path += "/" + file.Name + ".txt"
	err := os.WriteFile(path, []byte(file.Body), os.ModeDir)
	if err != nil {
		panic(err)
	}
}
