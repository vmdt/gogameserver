package domain

import "context"

type IRoomRepository interface {
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
	UpdateRoom(ctx context.Context, room *Room) (*Room, error)
	GetRoomByID(ctx context.Context, id string) (*Room, error)
}
