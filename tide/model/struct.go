package model

type File struct {
	Path string
	Body string
}
type AreaCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type URIinfo struct {
	URI  string
	Date string
	AreaCode
}
type InputDate struct {
	Year       int
	StartMonth int
	EndMonth   int
}
