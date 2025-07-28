package domain

import "time"

type User struct {
	ID           string `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string `gorm:"type:varchar(50);uniqueIndex" json:"username"`
	Email        string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password     string `gorm:"-" json:"-"`
	PasswordHash string `gorm:"not null" json:"-"`
	Nation       string `gorm:"type:varchar(50)" json:"nation"`

	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (u *User) TableName() string {
	return "identity_users"
}
