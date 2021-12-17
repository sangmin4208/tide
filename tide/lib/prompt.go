package lib

import (
	"fmt"

	"github.com/sangmin4208/tidy/model"
)

func Prompt() model.InputDate {

	var year int
	var sm int
	var em int
	for {
		fmt.Println("년도 시작(달) 종료(달) ")
		fmt.Scanln(&year, &sm, &em)
		if year != 0 && sm != 0 && em != 0 {
			break
		}
	}
	if sm < 1 || sm > 12 {
		sm = 1
	}
	if em < sm {
		em = sm
	}
	return model.InputDate{Year: year, StartMonth: sm, EndMonth: em}
}
