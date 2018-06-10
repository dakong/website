package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dakong/blog"
	"github.com/dakong/blog/mysql"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		panic(e)
	}
}

// Response to send back to the client side
func main() {
	db, err := mysql.Open("root:password@/blog")

	check(err)
	defer db.Close()

	blogService := mysql.InitializeService(db)
	blogRouter := blog.InitializeRouter()

	blogApp := &blog.App{Router: blogRouter, Service: blogService}
	appHandler := blogApp.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":3001", appHandler))
}
