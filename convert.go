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
	//"'고"가 들어있는것만 추출
	s = Filter(s, func (a string) bool {
		return strings.Contains(a, "고")
	})	

	times := make([]string, len(s))
	for i,v := range s {
		times[i] = trimmingTime(v)
	}
	info.Times = times
	return info
}

func Filter(s []string, f func(string) bool) []string{
	result := make([]string,0,len(s))

	for _,v := range s {
		if f(v){
			result = append(result, v)
		}
	}
	return result
}

// 시간데이터만 뽑는 용도
// 18:30/고/844 => 18:30
func trimmingTime (s string) string{
	return strings.Split(s,"/")[0]
}