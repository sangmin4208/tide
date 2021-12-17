package parse

import (
	"strings"

	"github.com/sangmin4208/filter/model"
)

func AtoTide(row string) model.Tide {
	temp := strings.Split(row, ", ")
	tide := model.Tide{
		Date: temp[0],
		Data: temp[1:],
	}
	return tide
}

func TidetoA(tide model.Tide) string {
	s := ""
	for _, data := range tide.Data {
		s += data + " "
	}
	return tide.Date + " " + s
}
