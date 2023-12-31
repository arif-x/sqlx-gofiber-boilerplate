package dashboard

import (
	"context"
	"fmt"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/google/uuid"
)

type PostCategoryRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.PostCategory, int, error)
	Show(UUID string) (model.PostCategoryShow, error)
	Store(model *model.StorePostCategory) (model.PostCategory, error)
	Update(UUID string, request *model.UpdatePostCategory) (model.PostCategory, error)
	Destroy(UUID string) (model.PostCategory, error)
}

type PostCategoryRepo struct {
	db *database.DB
}

func (repo *PostCategoryRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.PostCategory, int, error) {
	_select := "uuid, name, created_at, updated_at, deleted_at"
	_conditions := database.Search([]string{"name"}, search, "post_categories.deleted_at")
	_order := database.OrderBy(sort_by, sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM post_categories %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM post_categories %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.PostCategory
	for rows.Next() {
		var i model.PostCategory
		if err := rows.Scan(
			&i.UUID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, 0, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, 0, err
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, count, nil
}

func (repo *PostCategoryRepo) Show(UUID string) (model.PostCategoryShow, error) {
	var postCategory model.PostCategoryShow
	query := "SELECT uuid, name, created_at, updated_at, deleted_at FROM post_categories WHERE uuid = $1 AND post_categories.deleted_at IS NULL LIMIT 1"
	err := repo.db.QueryRowContext(context.Background(), query, UUID).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
		&postCategory.DeletedAt,
	)
	if err != nil {
		return model.PostCategoryShow{}, err
	}
	return postCategory, err
}

func (repo *PostCategoryRepo) Store(request *model.StorePostCategory) (model.PostCategory, error) {
	query := `INSERT INTO "post_categories" (uuid, name, created_at) VALUES($1, $2, $3) 
	RETURNING uuid, name, created_at`
	var postCategory model.PostCategory
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, time.Now()).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
	)
	if err != nil {
		return model.PostCategory{}, err
	}
	return postCategory, err
}

func (repo *PostCategoryRepo) Update(UUID string, request *model.UpdatePostCategory) (model.PostCategory, error) {
	query := `UPDATE "post_categories" SET name = $2, updated_at = $3 WHERE uuid = $1 
	RETURNING uuid, name, created_at, updated_at`
	var postCategory model.PostCategory
	err := repo.db.QueryRowContext(context.Background(), query, UUID, request.Name, time.Now()).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
	)
	if err != nil {
		return model.PostCategory{}, err
	}
	return postCategory, err
}

func (repo *PostCategoryRepo) Destroy(UUID string) (model.PostCategory, error) {
	query := `UPDATE "post_categories" SET updated_at = $2, deleted_at = $3 WHERE uuid = $1 
	RETURNING uuid, name, created_at, updated_at, deleted_at`
	var postCategory model.PostCategory
	err := repo.db.QueryRowContext(context.Background(), query, UUID, time.Now(), time.Now()).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
		&postCategory.DeletedAt,
	)
	if err != nil {
		return model.PostCategory{}, err
	}
	return postCategory, err
}

func NewPostCategoryRepo(db *database.DB) PostCategoryRepository {
	return &PostCategoryRepo{db}
}
