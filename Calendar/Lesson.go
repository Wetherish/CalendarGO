package Calendar

type Lesson struct {
	Id        string `json:"id" db:"id"`
	Subject   string `json:"subject" db:"Subject"`
	TeacherId string `json:"TeacherId" db:"TeacherId"`
	StudentId string `json:"StudentId" db:"StudentId"`
	Place     string `json:"Place" db:"Place"`
	Date      string `json:"Date" db:"Date"`
}

// type Place []string
