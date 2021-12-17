package parse

import (
	"strconv"
	"strings"

	"github.com/sangmin4208/filter/config"
	"github.com/sangmin4208/filter/model"
)

func AtoCriteria(row string) model.Criteria {
	temp := strings.Split(format(row), " ")
	criteria := model.Criteria{}
	criteria.Region = temp[config.POSITION_OF_REGION]
	h, err := strconv.Atoi(temp[config.POSITION_OF_HEIGHT])
	if err != nil {
		panic(err)
	}
	criteria.Height = h
	return criteria
}

func format(row string) string {
	row = strings.ReplaceAll(row, ",", "")
	for key, value := range config.Mapping {
		row = strings.Replace(row, key, value, 1)
	}
	return row
}
