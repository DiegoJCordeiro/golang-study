package people

type Service struct {
	Repository Repository
}

func (service *Service) GetPeople(id int) (*People, error) {
	return service.Repository.GetPeople(id)
}
