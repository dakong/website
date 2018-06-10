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

// BlogPost will fetch a given blog post from the database
func (service *BlogService) BlogPost(id int) (blogPost *model.BlogPost, err error) {
	row := service.DB.QueryRow("SELECT id, title, body FROM blog_post WHERE id = ?", id)
	blogPost = &model.BlogPost{}
	err = row.Scan(&blogPost.ID, &blogPost.Title, &blogPost.Body)

	return
}

// DeleteBlogPost ...
func (service *BlogService) DeleteBlogPost(id int) (blogPost *model.BlogPost, err error) {
	return nil, nil
}

// EditBlogPost ...
func (service *BlogService) EditBlogPost(id int) (blogPost *model.BlogPost, err error) {
	return nil, nil
}

// CreateBlogPost ...
func (service *BlogService) CreateBlogPost(id int) (blogPost *model.BlogPost, err error) {
	return nil, nil
}

// getBlogPostByID := func(res http.ResponseWriter, req *http.Request, p simplerouter.Params) {
// 	id, _ := p.GetValue("id")
// 	postList := make(BlogPosts, 0)

// 	rows, err := db.Query("SELECT id, title, body FROM blog_post WHERE id = ?", id)
// 	check(err)

// 	for rows.Next() {
// 		blogPost := BlogPost{}

// 		err = rows.Scan(&blogPost.ID, &blogPost.Title, &blogPost.Body)
// 		check(err)

// 		postList = append(postList, blogPost)
// 	}

// 	response := Response{Data: postList}
// 	result, err := json.Marshal(response)
// 	res.Write([]byte(result))
// }
