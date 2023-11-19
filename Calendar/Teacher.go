package Calendar

type Teacher struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"Name" db:"Name"`
	Street string `json:"Street" db:"Street"` // optional
}
