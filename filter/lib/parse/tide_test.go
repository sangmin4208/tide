package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTidy(t *testing.T) {
	row := "2021-03-01, 05:10/고/627, 11:33/저/-9, 17:33/고/675, --:--/-/--/--"
	result := AtoTide(row)

	assert.Equal(t, "2021-03-01", result.Date)
	assert.Equal(t, "05:10/고/627", result.Data[0])
	assert.Equal(t, "11:33/저/-9", result.Data[1])
	assert.Equal(t, "17:33/고/675", result.Data[2])
	assert.Equal(t, "--:--/-/--/--", result.Data[3])
}
