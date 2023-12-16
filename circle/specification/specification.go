package specification

import "ddd-demo/circle/entity"

type CircleSpecification struct{}

const maxMembersCount = 30

func (cs *CircleSpecification) IsFull(circle *entity.Circle) bool {
	return circle.CountMembers() > maxMembersCount
}
