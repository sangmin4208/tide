package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetInfo(t *testing.T) {
	line := "2021-07-05, 10:59/저/38, 19:03/-/23, 1-:--/-/--/--, 2-:--/-/--/--"
	info := getInfo(line)
	expected := "--:--"
	actual := info.Times[0]
	assert.Equal(t,expected,actual,"Time[0]")

	expected = "--:--"
	actual = info.Times[1]
	assert.Equal(t,expected,actual,"Time[1]")
	
	//output
	//0
}