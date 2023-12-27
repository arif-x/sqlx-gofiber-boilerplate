package public

import (
	"database/sql"

	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/public"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/paginate"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// PublicPostIndex func gets all post.
// @Description Get all post.
// @Summary Get all post
// @Tags Public Post
// @Accept json
// @Produce json
// @Param page query integer false "Page no"
// @Param page_size query integer false "records per page"
// @Success 200 {object} response.PublicPostsResponse
// @Failure 400,403 {object} response.ErrorResponse "Error"
// @Router /api/v1/public/post [get]
func PostIndex(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.Index(limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

// PublicPostByCategory func gets post by category.
// @Description Get post by category.
// @Summary Get post by category
// @Tags Public Post
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} response.PublicPostsByCategoryResponse
// @Failure 400,403 {object} response.ErrorResponse "Error"
// @Router /api/v1/public/post/category/{id} [get]
func PostCategoryPost(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	UUID := c.Params("id")
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.PostCategoryPost(UUID, limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

// PublicPostByUser func gets post by user.
// @Description Get post by user.
// @Summary Get post by user
// @Tags Public Post
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.PublicPostsByUserResponse
// @Failure 400,403 {object} response.ErrorResponse "Error"
// @Router /api/v1/public/post/user/{id} [get]
func UserPost(c *fiber.Ctx) error {
	page, limit, search, sort_by, sort := paginate.Paginate(c)
	UUID := c.Params("id")
	repository := repo.NewPostRepo(database.GetDB())

	posts, count, err := repository.UserPost(UUID, limit, uint(limit*(page-1)), search, sort_by, sort)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Index(c, page, limit, count, posts)
}

// PublicPostShow func gets single post.
// @Description Get single post.
// @Summary Get single post
// @Tags Public Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} response.PostResponse
// @Failure 400,403 {object} response.ErrorResponse "Error"
// @Router /api/v1/public/post/{id} [get]
func PostShow(c *fiber.Ctx) error {
	ID := c.Params("id")

	repository := repo.NewPostRepo(database.GetDB())
	post, err := repository.Show(ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFound(c, err)
		} else {
			return response.InternalServerError(c, err)
		}
	}

	return response.Show(c, post)
}
