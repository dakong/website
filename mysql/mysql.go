package mysql

import (
	"database/sql"

	"github.com/dakong/blog/model"
	_ "github.com/go-sql-driver/mysql"
)

// BlogService for interacting with the MySql database
type BlogService struct {
	DB *sql.DB
}

// Open a database connection
func Open(dataSource string) (*sql.DB, error) {
	return sql.Open("mysql", dataSource)
}

// InitializeService a database connection
func InitializeService(db *sql.DB) *BlogService {
	return &BlogService{DB: db}
}

// Create ...
func (service *BlogService) Create(id int) (blogPost *model.BlogPost, err error) {
	return nil, nil
}

// BlogPost will fetch a given blog post from the database
func (service *BlogService) BlogPost(id int) (blogPost *model.BlogPost, err error) {
	row := service.DB.QueryRow("SELECT id, title, body FROM blog_post WHERE id = ?", id)
	blogPost = &model.BlogPost{}
	err = row.Scan(&blogPost.ID, &blogPost.Title, &blogPost.Body)

	return
}

// BlogPosts will return all blog posts in the database
func (service *BlogService) BlogPosts() (blogPosts *model.BlogPosts, err error) {
	posts := make(model.BlogPosts, 0)

	rows, err := service.DB.Query("SELECT id, title, body FROM blog_post")

	for rows.Next() {
		blogPost := model.BlogPost{}

		err = rows.Scan(&blogPost.ID, &blogPost.Title, &blogPost.Body)
		posts = append(posts, blogPost)
	}

	blogPosts = &posts

	return
}

// Delete ...
func (service *BlogService) Delete(id int) (err error) {
	_, err = service.DB.Exec("DELETE FROM blog_post WHERE id = ?", id)
	return
}

// Update ...
func (service *BlogService) Update(id int) (blogPost *model.BlogPost, err error) {
	return nil, nil
}
