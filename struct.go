package main

type AreaCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Tide struct {
	Area  string
	Infos []*TideInfo
}

type TideInfo struct {
	Date  string
	Times []string
}