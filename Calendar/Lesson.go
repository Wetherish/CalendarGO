package Calendar

import "time"

type Lesson struct {
	Id        string    `json:"id" db:"id"`
	Subject   string    `json:"subject" db:"Subject"`
	TeacherId string    `json:"TeacherId" db:"TeacherId"`
	StudentId string    `json:"StudentId" db:"StudentId"`
	Place     string    `json:"Place" db:"Place"`
	Date      string    `json:"Date" db:"Date"`
	DataStart time.Time `json:"dataStart" db:"dataStart"`
	DataEnd   time.Time `json:"dataEnd" db:"dataEnd"`
}

