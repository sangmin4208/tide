package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func writeFileWithTide(tide *Tide) {
	times := strings.Join(tide.times, " ")
	err := ioutil.WriteFile(OUTPUT_PATH+"/"+tide.area+".txt", []byte(times), 0644)
	if err != nil {
		log.Fatal(err)
	}
}