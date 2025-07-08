package domain

type Player struct {
	ID     string `gorm:"type:uuid;primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(100)" json:"name"`
	UserId string `gorm:"type:uuid;default:null" json:"user_id"`
}
