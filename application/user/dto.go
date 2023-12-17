package user

type DTO struct {
	Id        string
	FirstName string
	LastName  string
}

func NewDTO(id, firstName, lastName string) *DTO {
	return &DTO{id, firstName, lastName}
}
