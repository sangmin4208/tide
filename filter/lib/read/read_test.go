package read

import (
	"testing"

	"github.com/sangmin4208/filter/config"
	"github.com/stretchr/testify/assert"
)

func TestDir(t *testing.T) {
	assert := assert.New(t)
	_, err := WalkDir(config.INPUT_PATH)
	assert.NoError(err)
	// assert.Equal(temp[0], "")

}

func TestRegion(t *testing.T) {
	region := Region("1.인천.txt")
	assert.Equal(t, region, "인천")
}
