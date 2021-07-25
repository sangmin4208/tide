package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func format(tide *Tide) string{
	line := ""
	for _, v := range tide.Info {
		line += v.Date + ", "
		line += strings.Join(v.Times," ")
		line += "\n"
	}
	return line
}

func writeFileWithTide(tide *Tide) {
	stringToWrite:= format(tide)
	err := ioutil.WriteFile(OUTPUT_PATH+"/"+tide.Area+".txt", []byte(stringToWrite), 0644)
	if err != nil {
		log.Fatal(err)
	}
}