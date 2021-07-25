package main

type Tide struct {
	Area string
	Info []TideInfo
}
type TideInfo struct {
	Date  string
	Times []string
}