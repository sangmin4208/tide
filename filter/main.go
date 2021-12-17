package main

import (
	"os"

	"github.com/sangmin4208/filter/config"
	"github.com/sangmin4208/filter/lib/filter"
	"github.com/sangmin4208/filter/lib/parse"
	"github.com/sangmin4208/filter/lib/read"
	"github.com/sangmin4208/filter/lib/write"
	"github.com/sangmin4208/filter/model"
)

var ctm map[string]model.Criteria

func init() {
	//OUTPUT 폴더 만들기
	if _, err := os.Stat(config.OUTPUT_PATH); os.IsNotExist(err) {
		os.Mkdir(config.OUTPUT_PATH, 0755)
	} else {
		os.RemoveAll(config.OUTPUT_PATH)
		os.Mkdir(config.OUTPUT_PATH, 0755)
	}

	// m := make(map[string] model.File)
	// 기준 파일 읽기
	temp, err := read.File(config.CRITERIA_DATA_PATH)
	if err != nil {
		panic(err)
	}
	ctm = make(map[string]model.Criteria)
	for _, row := range temp {
		ct := parse.AtoCriteria(row)
		ctm[ct.Region] = ct
	}
}

func main() {
	// 받아온 자료 파일구조 저장
	files, err := read.WalkDir(config.INPUT_PATH)
	if err != nil {
		panic(err)
	}

	tides := []model.Tide{}
	// 자료 객체로 파싱
	for _, file := range files {
		body, err := read.File(file.Path)
		if err != nil {
			panic(err)
		}
		region := read.Region(file.Name())
		for _, row := range body {
			tide := parse.AtoTide(row)
			tide.Region = region
			c, ok := ctm[tide.Region]
			if !ok {
				continue
			}
			tide = filter.Tide(tide, c)
			if len(tide.Data) != 0 {
				tides = append(tides, tide)
			}
		}
	}

	tm := make(map[string][]model.Tide)
	// 파일 작성을위해 map으로 파싱
	for _, tide := range tides {
		tm[tide.Region] = append(tm[tide.Region], tide)
	}
	results := parse.File(tm)
	for _, file := range results {
		write.File(config.OUTPUT_PATH, *file)
	}
}
