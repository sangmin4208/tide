package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func format(tide *Tide) string{
	times := strings.Join(tide.times, "\n")
	return times
}

func writeFileWithTide(tide *Tide) {
	stringToWrite:= format(tide)
	err := ioutil.WriteFile(OUTPUT_PATH+"/"+tide.area+".txt", []byte(stringToWrite), 0644)
	if err != nil {
		log.Fatal(err)
	}
}