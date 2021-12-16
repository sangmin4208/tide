package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sangmin4208/tidy/config"
	"github.com/sangmin4208/tidy/lib"
	"github.com/sangmin4208/tidy/model"
)

func main() {
	codes := ReadJsonFile()
	os.RemoveAll(config.OUTPUT_PATH)
	lib.MakeDirectory(config.OUTPUT_PATH)
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
	//월별
	for _, code := range codes {
		SearchAndSave(code, year, sm, em)
	}
}

// s: 시작 달, e: 끝나는 달
func SearchAndSave(code model.AreaCode, year int, s int, e int) {
	if s < 1 {
		s = 1
	}
	if e < s {
		e = s
	}
	date := ""
	for i := s; i <= e; i++ {
		if i >= 10 {
			date = strconv.Itoa(year) + strconv.Itoa(i)
		} else {
			date = strconv.Itoa(year) + "0" + strconv.Itoa(i)
		}
		path := config.OUTPUT_PATH + "/" + date
		lib.MakeDirectory(path)
		data := GetData(code, date)
		data = DeleteLine(data, 5)
		tide := new(model.Tide)
		tide.Area = code.Name
		lib.ConvertTide(data, tide)
		lib.WriteFile(path+"/"+tide.Area+".txt", tide)
	}
}

func ReadJsonFile() []model.AreaCode {
	var areaCode []model.AreaCode
	f, err := os.ReadFile(config.CODE_FILE)
	if err != nil {
		log.Fatal("./code.json을 못찾음", err)
	}
	json.Unmarshal(f, &areaCode)
	return areaCode
}

//date 포맷 -> 202101
func GetData(code model.AreaCode, date string) string {
	URI := config.DOWNLOAD_URI
	URI = strings.Replace(URI, "{code}", code.Code, -1)
	URI = strings.Replace(URI, "{date}", date, -1)

	log.Println("지역: "+code.Name, " 날짜: "+date+" URI: "+URI)
	res, err := http.Get(URI)
	if err != nil {
		log.Panicln(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	return lib.Decode(data)
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
