package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func convertToTide(fileName string) *Tide {
	area := getArea(fileName)
	times := []string{}
	fo, err := os.Open(INPUT_PATH + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
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
		area:  area,
		times: times,
	}
}

func getTime(line string) string {
	times := []string{}
	// 2021-07-01, 00:33/고/108, 06:54/저/42, 13:18/고/99, 18:59/저/51
	lines := strings.Split(string(line), ", ")
	// 00:33/고/108, 06:54/저/42, 13:18/고/99, 18:59/저/51
	lines = lines[1:]
	// [00:33/고/108 13:18/고/99]
	lines = Filter(lines, func(s string) bool {
		return !strings.Contains(s, "저")
	})
	//[00:33 고 108]
	for _, v := range lines {
		// [00:33]
		times = append(times, strings.Split(v, "/")[0])
	}

	return strings.Join(times, " ")
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


