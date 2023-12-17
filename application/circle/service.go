package circle

import (
	"context"
	"ddd-demo/domain/circle"
	"ddd-demo/domain/user"
	"errors"
)

type ApplicationService struct {
	userRep   user.Repository
	cds       *circle.DomainService
	circleRep circle.CircleRepository
}

func NewApplicationService(userRep user.Repository, cds *circle.DomainService, circleRep circle.CircleRepository) *ApplicationService {
	return &ApplicationService{userRep, cds, circleRep}
}

func (c *ApplicationService) Create(cmd CrateCommand) error {
	// オーナー取得
	owner, err := c.userRep.FindById(context.TODO(), cmd.Id)
	if err != nil {
		return err
	}

	// サークル名
	cn, err := circle.NewCircleName(cmd.Name)
	if err != nil {
		return err
	}

	// サークル これはファクトリ関数で
	circle := circle.Create(cn, owner)

	// サークル名の重複確認
	if c.cds.Exists(circle) {
		return errors.New("already eixists circle name " + cmd.Name)
	}

	// 永続化
	return c.circleRep.Save(context.TODO(), circle)
}
