package user

import (
	"context"
)

type Service struct {
	r *Repository
}

func NewService(r *Repository) *Service {
	return &Service{r}
}

func (s *Service) Exists(user *User) bool {
	_, err := s.r.FindByName(context.Background(), user.name.firstName, user.name.lastName)
	return err == nil
}
