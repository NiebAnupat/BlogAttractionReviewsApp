package controller

import (
	"io"
	"os"

	_AuthService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/service"
	_BlogModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/model"
	_BlogPostService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/service"
	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BlogPostControllerImpl struct {
	blogPostService    _BlogPostService.BlogPostService
	authService        _AuthService.AuthService
	fileStorageService _fileStorageService.FileStorageService
}

// DeleteBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) DeleteBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FavoriteBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) FavoriteBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) GetAllBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetBlogPostByID implements BlogPostController.
func (b *BlogPostControllerImpl) GetBlogPostByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// LikeBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) LikeBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// PostNewBlog implements BlogPostController.
func (b *BlogPostControllerImpl) PostNewBlog(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})

	}

	token = token[7:]

	uid, err := b.authService.VerifyToken(token)
	if err != nil || uid == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	createReqForm := new(_BlogModel.BlogPostCreateFormReq)

	if err := c.BodyParser(createReqForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request",
			"error":   err.Error(),
		})
	}

	createReq := new(_BlogModel.BlogPostCreateReq)

	var thumbnailBuffer io.Reader
	var filename string
	thumbnail, err := c.FormFile("thumbnail")
	if err != nil {
		defaultThumbnail, err := os.Open("public/defaultBlogThumbnail.jpg")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Thumbnail is required",
				"error":   err.Error(),
			})
		}
		defer defaultThumbnail.Close()
		thumbnailBuffer = defaultThumbnail

		filename, err = b.fileStorageService.UploadFile(thumbnailBuffer, uuid.New().String()+".png")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload thumbnail",
				"error":   err.Error(),
			})
		}

	} else {
		thumbnailBuffer, err = thumbnail.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to open thumbnail",
				"error":   err.Error(),
			})
		}
		thumbnail.Filename = uuid.New().String() + ".png"
		filename, err = b.fileStorageService.UploadFile(thumbnailBuffer, thumbnail.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload thumbnail",
				"error":   err.Error(),
			})
		}
	}

	createReq.Title = createReqForm.Title
	createReq.Description = createReqForm.Description
	createReq.Thumbnail = filename
	createReq.AuthorID = uid

	newBlogPost, err := b.blogPostService.CreateBlogPost(createReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create blog post",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Blog post created successfully",
		"blog":    newBlogPost,
	})
}

// AddContentToBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) AddContentToBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UnfavoriteBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) UnfavoriteBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UnlikeBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) UnlikeBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewBlogPostController(blogPostService _BlogPostService.BlogPostService, authService _AuthService.AuthService, fileStorageService _fileStorageService.FileStorageService) BlogPostController {
	return &BlogPostControllerImpl{
		blogPostService:    blogPostService,
		authService:        authService,
		fileStorageService: fileStorageService,
	}
}
