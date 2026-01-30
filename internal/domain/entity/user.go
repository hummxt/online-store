package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username    string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	FirstName   string    `gorm:"type:varchar(100)" json:"first_name"`
	LastName    string    `gorm:"type:varchar(100)" json:"last_name"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number"`
	Address     string    `gorm:"type:text" json:"address"`
	Role        string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
