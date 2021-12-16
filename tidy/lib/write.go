package lib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sangmin4208/tidy/config"
	"github.com/sangmin4208/tidy/model"
)

//폴더를 만드는 함수
//name은 폴더명
func MakeDirectory(path string) string {
	err := os.Mkdir(path, os.ModeDir)
	if err != nil {
		err = os.RemoveAll(path)
		if err != nil {
			log.Fatal("폴더 만들기에 실패..")
		}
		err = os.MkdirAll(path, os.ModeDir)
		if err != nil {
			log.Fatal("폴더 만들기에 실패..")
		}
	}
	return path
}

func WriteFile(file model.File) {
	path := fmt.Sprintf("%v/%v", config.OUTPUT_PATH, file.Path)
	err := os.WriteFile(path, []byte(file.Body), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

//text에서 한줄 삭제
//line - 삭제할 갯수
func DeleteLine(s string, line int) string {
	text := strings.Trim(s, " ")
	for i := 0; i < line; i++ {
		idx := strings.Index(text, "\n")
		text = text[idx+1:]
	}
	return text
}
