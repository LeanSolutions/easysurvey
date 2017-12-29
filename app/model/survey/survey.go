package survey

// Survey information
type Survey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}


var survey1, survey2 Survey

// GetAll returns all the surveys
func GetAll() []Survey {

	survey1.ID = "1"
	survey1.Name = "Survey 1"

	survey2.ID = "2"
	survey2.Name = "Survey 2"

	surveys := []Survey{survey1, survey2}
	return surveys
}