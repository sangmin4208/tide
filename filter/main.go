package main

import (
	"github.com/sangmin4208/filter/config"
	"github.com/sangmin4208/filter/lib"
	"github.com/sangmin4208/filter/lib/parse"
	"github.com/sangmin4208/filter/model"
)

func main() {
	temp, err := lib.Read(config.CRITERIA_DATA_PATH)
	if err != nil {
		panic(err)
	}
	ctm := make(map[string]model.Criteria)
	for _, row := range temp {
		ct := parse.AtoCriteria(row)
		ctm[ct.Region] = ct
	}

}
