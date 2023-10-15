package Models

type Teacher struct {
	Id      int      `json:"id"`
	Name    string   `json:"Name"`
	Subject []string `json:"Subject"`
}
