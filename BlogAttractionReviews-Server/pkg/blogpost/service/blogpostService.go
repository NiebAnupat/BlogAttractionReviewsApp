package service

import (
	_blogPostModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/model"
)

type BlogPostService interface {
	CreateBlogPost(blogPostCreateReq *_blogPostModel.BlogPostCreateReq) (*_blogPostModel.BlogPost, error)
	AddContentToBlogPost(blogPostID string, blogContent *_blogPostModel.BlogContentCreateReq) (*_blogPostModel.BlogContent, error)
	GetBlogPostByID(id string) (*_blogPostModel.BlogPost, error)
	GetAllBlogPost() ([]_blogPostModel.BlogPost, error)
	GetAllBlogPostByAuthorID(authorID string) ([]_blogPostModel.BlogPost, error)
	DeleteBlogPost(id string) error

	LikeBlogPost(userID, blogPostID string) error
	FavoriteBlogPost(userID, blogPostID string) error

	UnlikeBlogPost(userID, blogPostID string) error
	UnfavoriteBlogPost(userID, blogPostID string) error
}
