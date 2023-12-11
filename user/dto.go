package user

type DTO struct {
	id        string
	firstName string
	lastName  string
}

func NewDTO(id, firstName, lastName string) *DTO {
	return &DTO{id, firstName, lastName}
}
