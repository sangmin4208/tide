package lib

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sangmin4208/tidy/model"
)

func ReadJsonFile(path string) []model.AreaCode {
	var areaCode []model.AreaCode
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("./code.json을 못찾음", err)
	}
	json.Unmarshal(f, &areaCode)
	return areaCode
}
