package Models

type Teacher struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	StudentList []Student `json:"StudentList"`
	Subject     string    `json:"Subject"`
}
