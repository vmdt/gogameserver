package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/room/domain"
)

type RoomRepositoryImp struct {
	db *RoomDbContext
}

func NewRoomRepository(db *RoomDbContext) domain.IRoomRepository {
	return &RoomRepositoryImp{
		db: db,
	}
}

func (r *RoomRepositoryImp) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	if err := r.db.GetModelDB(&domain.Room{}).Create(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RoomRepositoryImp) GetRoomByID(ctx context.Context, id string) (*domain.Room, error) {
	var room domain.Room
	if err := r.db.GetModelDB(&domain.Room{}).Where("id = ?", id).First(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepositoryImp) UpdateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	if err := r.db.GetModelDB(&domain.Room{}).Where("id = ?", room.ID).Updates(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}
