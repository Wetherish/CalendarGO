package Models

type Lesson struct {
	Id      int     `json:"id"`
	Subject string  `json:"subject"`
	Teacher Teacher `json:"teacher"`
	Student Student ` json:"student"`
}
