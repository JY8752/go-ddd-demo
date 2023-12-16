package appservice

import (
	"context"
	"ddd-demo/circle/command"
	"ddd-demo/circle/domainservice"
	"ddd-demo/circle/factory"
	"ddd-demo/circle/repository"
	"ddd-demo/circle/value"
	"ddd-demo/user"
	"errors"
)

type CircleApplicationService struct {
	userRep   *user.Repository
	cds       *domainservice.CircleDomainService
	circleRep repository.CircleRepository
}

func NewCircleApplicationService(userRep *user.Repository, cds *domainservice.CircleDomainService, circleRep repository.CircleRepository) *CircleApplicationService {
	return &CircleApplicationService{userRep, cds, circleRep}
}

func (c *CircleApplicationService) Create(cmd command.CrateCircle) error {
	// オーナー取得
	owner, err := c.userRep.FindById(context.TODO(), cmd.Id)
	if err != nil {
		return err
	}

	// サークル名
	cn, err := value.NewCircleName(cmd.Name)
	if err != nil {
		return err
	}

	// サークル これはファクトリ関数で
	circle := factory.CreateCircle(cn, owner)

	// サークル名の重複確認
	if c.cds.Exists(circle) {
		return errors.New("already eixists circle name " + cmd.Name)
	}

	// 永続化
	return c.circleRep.Save(context.TODO(), circle)
}
