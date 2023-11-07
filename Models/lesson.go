package Models

type Lesson struct {
	Id        string `json:"id"`
	Subject   string `json:"subject"`
	TeacherId int    `json:"teacherID"`
	Street    string `json:"street"`
	Day       string `json:"day"`
	Time      string `json:"Time"`
	StudentId int 	 `json:"studentId"`
}
