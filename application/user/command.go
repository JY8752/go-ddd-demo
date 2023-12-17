package user

type UpdateCommand struct {
	id        string
	firstName string
	lastName  string
}

func NewUpdateCommand(id, firstName, lastName string) *UpdateCommand {
	return &UpdateCommand{id, firstName, lastName}
}

type DeleteCommand struct {
	id string
}

func NewDeleteCommand(id string) *DeleteCommand {
	return &DeleteCommand{id}
}
