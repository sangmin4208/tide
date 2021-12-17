package model

type Tide struct {
	Region string
	Date   string
	Data   []string
}
type Tides []Tide

type TideInfo struct {
	Region string
	Tides
}
