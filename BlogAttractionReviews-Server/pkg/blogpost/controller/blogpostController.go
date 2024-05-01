package controller

import "github.com/gofiber/fiber/v2"

type BlogPostController interface {
	PostNewBlog(c *fiber.Ctx) error
	GetBlogPostByID(c *fiber.Ctx) error
	GetAllBlogPost(c *fiber.Ctx) error
	DeleteBlogPost(c *fiber.Ctx) error

	LikeBlogPost(c *fiber.Ctx) error
	FavoriteBlogPost(c *fiber.Ctx) error
	UnlikeBlogPost(c *fiber.Ctx) error
	UnfavoriteBlogPost(c *fiber.Ctx) error
}
