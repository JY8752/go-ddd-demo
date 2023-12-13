package application

import (
	"context"
	"ddd-demo/user"
	"errors"
)

type Service struct {
	repository *user.Repository
	service    *user.Service
}

func NewService(r *user.Repository, s *user.Service) *Service {
	return &Service{r, s}
}

func (a *Service) Create(ctx context.Context, firstName, lastName string) (*user.DTO, error) {
	// 名前は値オブジェクト
	name, err := user.NewFullName(firstName, lastName)
	if err != nil {
		return nil, err
	}

	// Userをドメインモデルで表現
	u := user.NewFromName(name)

	// ドメインサービスでユーザーの存在確認
	if a.service.Exists(u) {
		return nil, errors.New("already exist user " + u.String())
	}

	// ユーザーを永続化
	result, err := a.repository.CreateUser(ctx, u.Id().String(), u.Name().FirstName(), u.Name().LastName())
	if err != nil {
		return nil, err
	}

	return toDTO(result), nil
}

func (a *Service) UpdateName(ctx context.Context, cmd *UpdateCommand) (*user.DTO, error) {
	// ユーザー情報を永続層から再構築
	u, err := a.repository.FindById(ctx, cmd.id)
	if err != nil {
		return nil, err
	}

	// 新しい名前の値オブジェクト
	newName, err := user.NewFullName(cmd.firstName, cmd.lastName)
	if err != nil {
		return nil, err
	}

	// ユーザー名の変更
	u.ChangeName(newName)

	// ドメインサービスで新しいユーザー名の存在確認
	if a.service.Exists(u) {
		return nil, errors.New("already exist user " + u.String())
	}

	// 永続化
	result, err := a.repository.UpdateName(ctx, u.Id().String(), u.Name().FirstName(), u.Name().LastName())
	if err != nil {
		return nil, err
	}

	return toDTO(result), nil
}

func toDTO(entity *user.User) *user.DTO {
	return user.NewDTO(
		entity.Id().String(),
		entity.Name().FirstName(),
		entity.Name().LastName(),
	)
}
