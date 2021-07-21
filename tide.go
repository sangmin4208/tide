package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const INPUT_PATH = "./data"
const OUTPUT_PATH = "./output"

type Tide struct {
	area string
	times []string
}

func main() {
	tides := []*Tide{}
	fileNames := getFileNames(INPUT_PATH)
	for _,fn := range fileNames{
		 tide := getContent(fn)
		tides = append(tides, tide)
	}

	for _,tide := range tides{
		writeFileWithTide(tide)
	}
}

func writeFileWithTide(tide *Tide){
	times := strings.Join(tide.times," ")
	err := ioutil.WriteFile(OUTPUT_PATH+"/"+tide.area+".txt",[]byte(times),0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getContent(fileName string) *Tide {
	area := getArea(fileName)
	times := []string{}
	fo, err := os.Open(INPUT_PATH + "/" + fileName)
	defer fo.Close()
	if err != nil{
		log.Fatal(err)
	}
	reader := bufio.NewReader(fo)
	trimLine(reader, 5)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		  }
		 times = append(times, getTime(string(line)))
	}
	return &Tide{
		area: area,
		times:times,
	}
}

func getTime(line string) string{
	times:=[]string{}
	// 2021-07-01, 00:33/고/108, 06:54/저/42, 13:18/고/99, 18:59/저/51
	lines := strings.Split(string(line),", ")
	// 00:33/고/108, 06:54/저/42, 13:18/고/99, 18:59/저/51
	lines = lines[1:]
	// [00:33/고/108 13:18/고/99]
	lines = Filter(lines,func(s string) bool{
		return !strings.Contains(s,"저")
	})
	for _,v := range lines {
		///[00:33 고 108]
		times = append(times,strings.Split(v,"/")[0] )
	} 
	
	return strings.Join(times," ")
}

func getFileNames(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	names := make([]string, 0, len(files))
	for _, f := range files {
		names = append(names, f.Name())
	}
	return names
}

func getArea(fn string) (area string) {
	area = fn[21 : len(fn)-4]
	return
}

func trimLine(reader *bufio.Reader, n int) {
	for i := 0; i < n; i++ {
		reader.ReadLine()
	}
}


func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
	   if f(v) {
		  vsf = append(vsf, v)
	   }
	}
	return vsf
 }
