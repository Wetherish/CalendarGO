package Calendar

type Student struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"Name" db:"Name"`
	Street    string `json:"street" db:"street"` 
}
