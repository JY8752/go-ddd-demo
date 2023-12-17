package circle

type Specification struct{}

const maxMembersCount = 30

func (cs *Specification) IsFull(circle *Circle) bool {
	return circle.CountMembers() > maxMembersCount
}
