package Models

type Teacher struct {
	Id        int    `json:"id"`
	Name      string `json:"Name"`
	Street    string `json:"street"`  // optional
}
