package auth

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string    `json:"password" gorm:"type:varchar(255);not null"`
}
