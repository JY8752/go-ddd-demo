package application

import (
	"context"
	"ddd-demo/user"
)

type DeleteService struct {
	repository *user.Repository
}

func NewDeleteService(r *user.Repository) *DeleteService {
	return &DeleteService{r}
}

func (ds *DeleteService) Handle(ctx context.Context, cmd *DeleteCommand) error {
	_, err := ds.repository.Delete(ctx, cmd.id)
	return err
}
