package domain

type IBattleShipRepository interface {
	CreateBoard(*BattleShip) (*BattleShip, error)
	GetBoardGameByPlayerId(playerId string, roomId string) (*BattleShip, error)
}
