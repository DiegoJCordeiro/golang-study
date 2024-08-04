package people

type Repository interface {
	GetPeople(id int) (*People, error)
}
