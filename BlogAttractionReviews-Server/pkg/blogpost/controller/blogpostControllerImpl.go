package controller

import (
	"io"
	"os"
	"strconv"

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
	blogID := c.Params("bid")
	err := b.blogPostService.DeleteBlogPost(blogID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete blog post",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Blog post deleted successfully",
	})
}

// FavoriteBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) FavoriteBlogPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllBlogPost implements BlogPostController.
func (b *BlogPostControllerImpl) GetAllBlogPost(c *fiber.Ctx) error {

	authorID := c.Query("aid")
	if authorID != "" {
		blogPosts, err := b.blogPostService.GetAllBlogPostByAuthorID(authorID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get blog posts by author",
				"error":   err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Get blog posts by author successfully",
			"blogs":   blogPosts,
		})
	} else {
		blogPosts, err := b.blogPostService.GetAllBlogPost()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get all blog posts",
				"error":   err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Get all blog posts successfully",
			"blogs":   blogPosts,
		})
	}

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

	userModel, err := b.authService.VerifyToken(token)
	if err != nil {
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
	createReq.AuthorID = userModel.ID

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
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	token = token[7:]
	_, err := b.authService.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	order, err := strconv.Atoi("1")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Order is invalid",
			"error":   err.Error(),
		})
	}
	contentType, err := strconv.Atoi(c.FormValue("type"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Type is invalid",
			"error":   err.Error(),
		})
	}
	// var value string
	value := make(chan string, 1)

	go func() {
		if contentType == _BlogModel.BlogContentText {
			value <- c.FormValue("value")
		} else if contentType == _BlogModel.BlogContentImage {
			file, err := c.FormFile("value")
			if err != nil {
				c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Value is required",
					"error":   err.Error(),
				})
				return
			}
			fileBuffer, err := file.Open()
			if err != nil {
				c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Failed to open file",
					"error":   err.Error(),
				})
				return
			}
			defer fileBuffer.Close()
			filename := uuid.New().String() + ".png"
			filename, err = b.fileStorageService.UploadFile(fileBuffer, filename)
			if err != nil {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Failed to upload file",
					"error":   err.Error(),
				})
				return
			}
			value <- filename
		} else {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid content type",
			})
			return

		}
	}()

	createReq := new(_BlogModel.BlogContentCreateReq)
	createReq.BlogID = c.FormValue("blogID")
	createReq.Order = order
	createReq.Type = contentType
	createReq.Value = <-value

	close(value)

	newContent, err := b.blogPostService.AddContentToBlogPost(createReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add content to blog post",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Content added to blog post successfully",
		"content": newContent,
	})

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
