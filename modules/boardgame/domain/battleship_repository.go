package domain

type IBattleShipRepository interface {
	AddOrUpdate(*BattleShip) (*BattleShip, error)
	GetBoardGameByPlayerId(playerId string, roomId string) (*BattleShip, error)
}
