package employee

type Employee struct {
	Id      int
	Name    string
	Title   string
	Company string
}

func (e *Employee) Create() (err error) {
	return nil
}
