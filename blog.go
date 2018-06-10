package blog

import (
	"encoding/json"
	"net/http"

	"github.com/dakong/blog/mysql"
	"github.com/dakong/simplerouter"
)

// App ...
type App struct {
	Router  *simplerouter.Router
	Service *mysql.BlogService
}

// Response to send back to the client side
type Response struct {
	Data interface{}
}

// AppError detailing what went wrong in the application
type AppError struct {
	Error   error
	Message string
	Code    int
}

type resourceHandler func(req *http.Request, p simplerouter.Params) (res *Response, err *AppError)

func errorHandler(handler resourceHandler) simplerouter.Handle {
	fn := func(res http.ResponseWriter, req *http.Request, p simplerouter.Params) {
		response, err := handler(req, p)
		if err != nil {
			res.WriteHeader(err.Code)
			res.Write(toJSON(err))
		} else {
			res.Write(toJSON(response))
		}
	}
	return fn
}

func toJSON(w interface{}) []byte {
	result, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	return result
}

// InitializeRouter ...
func InitializeRouter() *simplerouter.Router {
	return simplerouter.New()
}

// RegisterRoutes will register the CRUD routes for the blog application
func (app *App) RegisterRoutes() *simplerouter.Router {

	// Define our CRUD rest API
	// TODOD: Currently there's a bug in the simple router
	app.Router.Post("/blog", errorHandler(app.CreateBlogPost))
	app.Router.Get("/blog/:id", errorHandler(app.ReadBlogPost))
	app.Router.Get("/blog", errorHandler(app.ReadAllBlogPosts))
	app.Router.Put("/blog/:id", errorHandler(app.UpdatBlogPost))
	app.Router.Delete("/blog/:id", errorHandler(app.DeleteBlogPost))

	return app.Router
}
