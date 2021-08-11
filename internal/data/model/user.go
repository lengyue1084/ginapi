package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	*gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
