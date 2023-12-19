package dashboard

import (
	"time"

	"github.com/google/uuid"
)

type PostCategory struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type PostCategoryShow struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type StorePostCategory struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type UpdatePostCategory struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
