package filter

import (
	"strconv"
	"strings"

	"github.com/sangmin4208/filter/model"
)

func Tide(tide model.Tide, citeria model.Criteria) model.Tide {
	newData := []string{}
	for _, data := range tide.Data {
		if tideHelper(data, citeria.Height) {
			newData = append(newData, data)
		}
	}
	tide.Data = newData
	return tide
}

func tideHelper(data string, criteria int) bool {
	height, err := strconv.Atoi(strings.Split(data, "/")[2])
	//"--" 숫자가 아니면 에러 발생
	if err != nil {
		return false
	}
	if criteria >= height {
		return false
	}
	return true
}
