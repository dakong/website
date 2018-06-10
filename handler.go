package blog

import (
	"net/http"
	"strconv"

	"github.com/dakong/simplerouter"
)

// CreateBlogPost ...
func (app *App) CreateBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	return &Response{Data: "Create a blog post"}, nil
}

// DeleteBlogPost ...
func (app *App) DeleteBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	return &Response{Data: "Delete a blog post"}, nil
}

// EditBlogPost ...
func (app *App) EditBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	return &Response{Data: "Edit a blog post"}, nil
}

// GetBlogPost ...
func (app *App) GetBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	val, _ := p.GetValue("id")
	id, _ := strconv.Atoi(val)

	blogPost, err := app.Service.BlogPost(id)

	if err != nil {
		return nil, &AppError{Error: err, Message: "Error getting blog post", Code: 500}
	}

	return &Response{Data: blogPost}, nil
}
