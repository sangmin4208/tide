package main

import (
	"golang.org/x/text/encoding/korean"
)

func Decode(p []byte) string {
	t, _ := korean.EUCKR.NewDecoder().Bytes(p)
	return string(t)
}
