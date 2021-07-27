package main

import (
	"strings"
)

//하나의 텍스트파일로 Tide 데이터를 채움
func ConvertTide(text string, tide *Tide){
	infos := []*TideInfo{}
	for {
		idx := strings.Index(text,"\n")
		if idx == -1{
			break
		}
		line := text[:idx]
		text = text[idx+1:]
		info := getInfo(line)
		infos = append(infos, info)
	}
	tide.Infos = infos
}

//한줄로 TideInfo를 만들어서 리턴
func getInfo(line string) *TideInfo{
	info := new(TideInfo)
	s := strings.Split(line, ",")
	info.Date = s[0]
	s = s[1:]

	times := getTimes(s)
	info.Times = times
	return info
}


// 시간 데이터만 빼옴
// 18:30/고/844 => 18:30
func trimmingTime (s string) string{
	time := strings.Split(s,"/")[0]
	return strings.Trim(time," ")
} 

//16:16/고/41, --:--/-/--/--, --:--/-/--/--, --:--/-/--/-- => 16:16, --:--
//날짜 뺀 한줄을 받아서 시간 뽑아오기
func getTimes (line []string) []string{
	timeIdx := []int{-1,-1}
	for i,v := range line {
		if strings.Contains(v,"고") {
			if i == 0 || i == 2 {
				timeIdx[0] = 0
				timeIdx[1] = 2
			}else{
				timeIdx[0] = 1
				timeIdx[1] = 3
			}
			break
		}
	}
	if timeIdx[0] == -1 {
		return []string{"--:--","--:--"}
	}
	time1 := trimmingTime(line[timeIdx[0]])
	time2 := trimmingTime(line[timeIdx[1]])


	return []string{time1,time2}
}