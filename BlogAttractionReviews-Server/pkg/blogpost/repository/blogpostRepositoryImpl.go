package repository

import (
	"log"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	_blogPostException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/exception"
)

type BlogPostRepositoryImpl struct {
	db database.Database
}

// Create implements BlogPostRepository.
func (b *BlogPostRepositoryImpl) Create(blogPostEntity *entities.BlogPost) (*entities.BlogPost, error) {
	err := b.db.Connect().Create(blogPostEntity).Error
	if err != nil {
		log.Println(err)
		return nil, &_blogPostException.BlogPostCreate{}
	}
	return blogPostEntity, nil
}

// CreateContent implements BlogPostRepository.
func (b *BlogPostRepositoryImpl) CreateContent(blogContentEntity *entities.BlogContent) (*entities.BlogContent, error) {
	err := b.db.Connect().Create(blogContentEntity).Error
	if err != nil {
		return nil, &_blogPostException.BlogContentCreate{}
	}
	return blogContentEntity, nil
}

// Delete implements BlogPostRepository.
func (b *BlogPostRepositoryImpl) Delete(id string) error {
	blogPost := &entities.BlogPost{}
	err := b.db.Connect().First(blogPost, "id = ?", id).Error
	if err != nil {
		return &_blogPostException.BlogPostNotFound{ID: id}
	}

	err = b.db.Connect().Delete(blogPost).Error
	if err != nil {
		return &_blogPostException.BlogPostDelete{}
	}
	return nil
}

// FindAll implements BlogPostRepository.
func (b *BlogPostRepositoryImpl) FindAll() ([]entities.BlogPost, error) {
	blogPosts := []entities.BlogPost{}
	err := b.db.Connect().Find(&blogPosts).Error
	if err != nil {
		return nil, &_blogPostException.BlogPostNotFound{}
	}

	// loop for get content for each blog post
	for i, blogPost := range blogPosts {
		blogContent := []*entities.BlogContent{}
		err := b.db.Connect().Find(&blogContent, "blog_id = ?", blogPost.ID).Error
		if err != nil {
			return nil, err
		}
		blogPosts[i].Contents = blogContent
	}

	return blogPosts, nil
}

// FindByID implements BlogPostRepository.
func (b *BlogPostRepositoryImpl) FindByID(id string) (*entities.BlogPost, error) {
	blogPost := &entities.BlogPost{}
	err := b.db.Connect().First(blogPost, "id = ?", id).Error
	if err != nil {
		return nil, &_blogPostException.BlogPostNotFound{ID: id}
	}

	blogContent := []*entities.BlogContent{}
	err = b.db.Connect().Find(&blogContent, "blog_id = ?", blogPost.ID).Error
	if err != nil {
		return nil, err
	}

	blogPost.Contents = blogContent

	return blogPost, nil
}

// FindAllByAuthorID implements BlogPostRepository.
func (b *BlogPostRepositoryImpl) FindAllByAuthorID(authorID string) ([]entities.BlogPost, error) {
	blogPosts := []entities.BlogPost{}
	err := b.db.Connect().Find(&blogPosts, "author_id = ?", authorID).Error
	if err != nil {
		return nil, &_blogPostException.BlogPostNotFound{}
	}

	// loop for get content for each blog post
	for i, blogPost := range blogPosts {
		blogContent := []*entities.BlogContent{}
		err := b.db.Connect().Find(&blogContent, "blog_id = ?", blogPost.ID).Error
		if err != nil {
			return nil, err
		}
		blogPosts[i].Contents = blogContent
	}

	return blogPosts, nil
}

func NewBlogPostRepositoryImpl(db database.Database) BlogPostRepository {
	return &BlogPostRepositoryImpl{db: db}
}
