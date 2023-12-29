package dashboard

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	UUID      uuid.UUID  `db:"uuid" json:"uuid"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type ShowRole struct {
	UUID      uuid.UUID  `db:"uuid" json:"uuid"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type StoreRole struct {
	Name string `json:"name" form:"name"`
}

type UpdateRole struct {
	Name string `json:"name" form:"name"`
}
