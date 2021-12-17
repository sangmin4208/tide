package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/sangmin4208/tidy/model"
)

// func GetData(code model.AreaCode, date string) string {
// 	URI := config.DOWNLOAD_URI
// 	URI = strings.Replace(URI, "{code}", code.Code, -1)
// 	URI = strings.Replace(URI, "{date}", date, -1)

//
// 	res, err := http.Get(URI)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	data, _ := ioutil.ReadAll(res.Body)
// 	return lib.Decode(data)
// }

func GetData(URI string) string {
	res, err := http.Get(URI)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("URI: " + URI)
	data, _ := ioutil.ReadAll(res.Body)
	return Decode(data)
}

func GetURI(base string, code model.AreaCode, year, month int) string {
	URI := strings.ReplaceAll(base, "{code}", code.Code)
	return strings.ReplaceAll(URI, "{date}", dateFormat(year, month))
}

func dateFormat(year int, month int) string {
	date := ""
	if month >= 10 {
		date = fmt.Sprintf("%v%v", year, month)
	} else {
		date = fmt.Sprintf("%v0%v", year, month)
	}
	return date
}
