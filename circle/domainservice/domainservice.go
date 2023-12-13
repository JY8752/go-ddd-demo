package domainservice

import (
	"context"
	"ddd-demo/circle/entity"
	"ddd-demo/circle/repository"
)

type CircleDomainService struct {
	circleRep repository.CircleRepository
}

func NewCircleDomainService(circleRep repository.CircleRepository) *CircleDomainService {
	return &CircleDomainService{circleRep}
}

func (c *CircleDomainService) Exists(circle *entity.Circle) bool {
	_, err := c.circleRep.FindByName(context.TODO(), circle.Name)
	return err == nil
}
