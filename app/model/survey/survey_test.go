package survey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllReturnsAllSurveys(t *testing.T) {
	assert := assert.New(t)
	surveys := GetAll()

	assert.NotZero(surveys, "It should return more than one survey")
}
