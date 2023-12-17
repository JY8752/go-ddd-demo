package circle

import (
	"context"
)

type DomainService struct {
	circleRep CircleRepository
}

func NewDomainService(circleRep CircleRepository) *DomainService {
	return &DomainService{circleRep}
}

func (c *DomainService) Exists(circle *Circle) bool {
	_, err := c.circleRep.FindByName(context.TODO(), circle.Notify().Name)
	return err == nil
}
