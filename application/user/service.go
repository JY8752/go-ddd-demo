package user

import (
	"context"
	"ddd-demo/domain/user"
	"errors"
)

type ApplicationService struct {
	repository user.Repository
	service    *user.Service
}

func NewApplicationService(r user.Repository, s *user.Service) *ApplicationService {
	return &ApplicationService{r, s}
}

func (a *ApplicationService) Create(ctx context.Context, firstName, lastName string) (*DTO, error) {
	// 名前は値オブジェクト
	name, err := user.NewFullName(firstName, lastName)
	if err != nil {
		return nil, err
	}

	// Userをドメインモデルで表現
	u := user.Create(name)

	// ドメインサービスでユーザーの存在確認
	if a.service.Exists(u) {
		return nil, errors.New("already exist user " + u.String())
	}

	// ユーザーを永続化
	result, err := a.repository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return toDTO(result), nil
}

func (a *ApplicationService) UpdateName(ctx context.Context, cmd *UpdateCommand) (*DTO, error) {
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
	result, err := a.repository.UpdateName(ctx, u)
	if err != nil {
		return nil, err
	}

	return toDTO(result), nil
}

func toDTO(entity *user.User) *DTO {
	notification := entity.Notify()
	return NewDTO(
		notification.Id.String(),
		notification.Name.FirstName(),
		notification.Name.LastName(),
	)
}
