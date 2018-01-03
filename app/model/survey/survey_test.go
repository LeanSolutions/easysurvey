package survey

import (
	"testing"
)

func TestGetAllReturnsAllSurveys(t *testing.T) {
	surveys := GetAll()

	if len(surveys) <= 0 {
		t.Errorf("Expected more than one survey to be returned, but got %d instead.", len(surveys))
	}
}