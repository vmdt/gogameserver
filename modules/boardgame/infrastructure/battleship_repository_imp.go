package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
)

type BattleShipRepositoryImp struct {
	db      *BoardGameDbContext
	generic *postgresgorm.GenericRepository[domain.BattleShip]
	ctx     context.Context
}

func NewBattleShipRepositoryImp(db *BoardGameDbContext) domain.IBattleShipRepository {
	repo := postgresgorm.NewGenericRepository[domain.BattleShip](db.GetModelDB(&domain.BattleShip{}))
	return &BattleShipRepositoryImp{
		db:      db,
		generic: repo,
	}
}

func (repo *BattleShipRepositoryImp) CreateBoard(battleShip *domain.BattleShip) (*domain.BattleShip, error) {
	if err := repo.generic.Add(battleShip, repo.ctx); err != nil {
		return nil, err
	}
	return battleShip, nil
}

func (repo *BattleShipRepositoryImp) GetBoardGameByPlayerId(playerId string, roomId string) (*domain.BattleShip, error) {
	var params = domain.BattleShip{
		PlayerId: uuid.MustParse(playerId),
		RoomId:   uuid.MustParse(roomId),
	}
	return repo.generic.Get(&params, repo.ctx), nil
}
