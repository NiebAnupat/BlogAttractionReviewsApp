package service

import (
	"log"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	_blogPostException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/exception"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/model"
	_blogPostModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/model"
	_blogPostRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/repository"
	"github.com/google/uuid"
)

type BlogPostServiceImpl struct {
	blogPostRepository _blogPostRepository.BlogPostRepository
}

// CreateBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) CreateBlogPost(blogPostCreateReq *model.BlogPostCreateReq) (*_blogPostModel.BlogPost, error) {
	blogID := uuid.New().String()

	blogPost := &entities.BlogPost{
		ID:          blogID,
		Title:       blogPostCreateReq.Title,
		Description: blogPostCreateReq.Description,
		Thumbnail:   blogPostCreateReq.Thumbnail,
		AuthorID:    blogPostCreateReq.AuthorID,
	}

	newBlogPost, err := b.blogPostRepository.Create(blogPost)
	if err != nil {
		return nil, &_blogPostException.BlogPostCreate{}
	}

	return newBlogPost.ToBlogPostModel(), nil

}

// AddContentToBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) AddContentToBlogPost(blogContent *_blogPostModel.BlogContentCreateReq) (*_blogPostModel.BlogContent, error) {
	content := new(entities.BlogContent)
	content.ID = uuid.New().String()
	content.BlogID = blogContent.BlogID
	content.Order = blogContent.Order
	content.Type = blogContent.Type
	content.Value = blogContent.Value
	log.Println(content)
	newContent, err := b.blogPostRepository.CreateContent(content)
	if err != nil {
		return nil, &_blogPostException.BlogContentCreate{}
	}

	return newContent.ToBlogContentModel(), nil
}

// DeleteBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) DeleteBlogPost(id string) error {
	panic("unimplemented")
}

// FavoriteBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) FavoriteBlogPost(userID string, blogPostID string) error {
	panic("unimplemented")
}

// GetAllBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) GetAllBlogPost() ([]model.BlogPost, error) {
	panic("unimplemented")
}

// GetAllBlogPostByAuthorID implements BlogPostService.
func (b *BlogPostServiceImpl) GetAllBlogPostByAuthorID(authorID string) ([]model.BlogPost, error) {
	panic("unimplemented")
}

// GetBlogPostByID implements BlogPostService.
func (b *BlogPostServiceImpl) GetBlogPostByID(id string) (*model.BlogPost, error) {
	panic("unimplemented")
}

// LikeBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) LikeBlogPost(userID string, blogPostID string) error {
	panic("unimplemented")
}

// UnfavoriteBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) UnfavoriteBlogPost(userID string, blogPostID string) error {
	panic("unimplemented")
}

// UnlikeBlogPost implements BlogPostService.
func (b *BlogPostServiceImpl) UnlikeBlogPost(userID string, blogPostID string) error {
	panic("unimplemented")
}

func NewBlogPostServiceImpl(blogPostRepository _blogPostRepository.BlogPostRepository) BlogPostService {
	return &BlogPostServiceImpl{
		blogPostRepository: blogPostRepository,
	}
}
