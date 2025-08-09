package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/pkg/cryptography/hasher"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string    `gorm:"type:varchar(50)" json:"username"` // display name
	Email        string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password     string    `gorm:"-" json:"-"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Nation       string    `gorm:"type:varchar(50)" json:"nation"`
	IsSso        bool      `gorm:"default:false" json:"is_sso"`
	Provider     string    `gorm:"type:varchar(50)" json:"provider"`

	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (u *User) TableName() string {
	return "identity_users"
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	if u.PasswordHash == "" {
		u.GenPassv2()
	}

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	if u.Password != "" {
		u.GenPassv2()
	}
	return
}

func (u *User) GenPassv2() {
	u.PasswordHash, _ = hasher.GenerateFromPassword(u.Password)
}

func (u *User) ValidatePassword(password string) (bool, error) {
	return hasher.ComparePasswordAndHash(password, u.PasswordHash)
}

func (u *User) ChangePassword(password string, newPassword string) error {
	result, err := hasher.ComparePasswordAndHash(password, u.PasswordHash)
	if err == nil && result {
		u.Password = newPassword
	}
	return err
}
