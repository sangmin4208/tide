package main

import (
	"log"
	"os"
	"strings"
)

//폴더를 만드는 함수
//name은 폴더명
func MakeDirectory(path string) string{
	err := os.MkdirAll(path, os.ModeDir)
	if err != nil{
		log.Fatal("폴더 만들기에 실패..")
	}
	return path
}


func formatline(info *TideInfo) string{
	line := info.Date + ","
	line += strings.Join(info.Times," ")
	line += "\n"
	return line
}

func WriteFile(path string, tide *Tide){
	lines := make([]string, len(tide.Infos))
	for i,v := range tide.Infos{
		lines[i] = formatline(v)
	}
	w := []byte(strings.Join(lines,""))
	os.WriteFile(path ,w, os.ModePerm)
}