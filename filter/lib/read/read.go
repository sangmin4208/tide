package read

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sangmin4208/filter/model"
)

func File(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	rows := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}
	return rows, nil
}

func WalkDir(root string) ([]model.FileInfo, error) {
	files := []model.FileInfo{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, err := d.Info()

		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, model.FileInfo{path, info})
		}
		return nil
	})
	return files, err
}

func Region(filename string) string {
	return strings.Split(filename, ".")[1]
}
