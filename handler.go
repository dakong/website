package blog

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dakong/simplerouter"
)

// CreateBlogPost ...
func (app *App) CreateBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	return &Response{Data: "Create a blog post"}, nil
}

// ReadBlogPost ...
func (app *App) ReadBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	val, _ := p.GetValue("id")
	id, _ := strconv.Atoi(val)

	blogPost, err := app.Service.BlogPost(id)

	if err != nil {
		appError := &AppError{Error: err, Message: "Error getting blog post", Code: http.StatusInternalServerError}

		if err == sql.ErrNoRows {
			appError.Message = "Could not find blog post"
			appError.Code = http.StatusNotFound
		}

		return nil, appError
	}

	return &Response{Data: blogPost}, nil
}

// ReadAllBlogPosts ...
func (app *App) ReadAllBlogPosts(req *http.Request, p simplerouter.Params) (*Response, *AppError) {

	blogPosts, err := app.Service.BlogPosts()

	if err != nil {
		return nil, &AppError{Error: err, Message: "Error getting all blog posts", Code: http.StatusInternalServerError}
	}

	return &Response{Data: blogPosts}, nil
}

// UpdatBlogPost ...
func (app *App) UpdatBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	return &Response{Data: "Edit a blog post"}, nil
}

// DeleteBlogPost ...
func (app *App) DeleteBlogPost(req *http.Request, p simplerouter.Params) (*Response, *AppError) {
	val, _ := p.GetValue("id")
	id, _ := strconv.Atoi(val)

	err := app.Service.Delete(id)

	if err != nil {
		appError := &AppError{Error: err, Message: "Error deleting blog post", Code: http.StatusInternalServerError}

		if err == sql.ErrNoRows {
			appError.Message = "Could not find blog post"
			appError.Code = http.StatusNotFound
		}

		return nil, appError
	}

	return &Response{Data: ""}, nil

}
