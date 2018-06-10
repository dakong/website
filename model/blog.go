package model

import (
	"strconv"
)

// BlogPost is the struture for each blog post
type BlogPost struct {
	ID    int
	Title string
	Body  string
}

func (post *BlogPost) String() (result string) {
	id := strconv.Itoa(post.ID)
	result = "{\n	ID: " + id +
		"\n	Title: " + post.Title +
		"\n	Body: " + post.Body + "\n}"

	return
}

// BlogPosts are a list of blog posts
type BlogPosts []BlogPost
