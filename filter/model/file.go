package model

import "io/fs"

type File struct {
	Name string
	Body string
}

type FileInfo struct {
	Path string
	fs.FileInfo
}
