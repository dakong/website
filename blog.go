package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dakong/blog/models"
	"github.com/google/uuid"
)

const BlogPostDirPath = "./posts/"

// Map string representation of a map to an int representation
var Month_Map = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		panic(e)
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// Get all the blog post files in the blog post directory
	blogPosts, err := ioutil.ReadDir(BlogPostDirPath)
	check(err)

	results := make([]models.Post, 0)

	// Iterate through each file
	for _, post := range blogPosts {
		var newPost models.Post

		file := (BlogPostDirPath + post.Name())
		fileContents, err := ioutil.ReadFile(file)
		check(err)

		// Convert from text to our Post Struct
		err = json.Unmarshal(fileContents, &newPost)
		check(err)

		results = append(results, newPost)
	}

	postsJson, err := json.Marshal(results)
	w.WriteHeader(http.StatusOK)
	w.Write(postsJson)

}

func convertSpaces(s string) string {
	strArray := strings.Split(s, " ")
	return strings.Join(strArray, "_")
}

func savePostToFile(post models.Post) {
	// Format the current datetime for the filename
	currentDate := time.Now()

	// Create elements for the filename to save the blog post
	yearString := strconv.Itoa(currentDate.Year())
	monthString := strconv.Itoa(Month_Map[currentDate.Month().String()])
	dayString := strconv.Itoa(currentDate.Day())
	titleString := convertSpaces(post.Title)

	// Join all file elements to build to filename
	fileNameElements := []string{yearString, monthString, dayString, titleString}
	fileName := strings.Join(fileNameElements, "_")
	fileName = BlogPostDirPath + fileName + ".json"

	// Populate the rest of the Blog Post data object
	post.PostedDate = currentDate.Format(time.UnixDate)
	post.Id = uuid.New().String()

	// Convert struct into a string to store into JSON file
	fileContents, err := json.MarshalIndent(post, "", " ")
	check(err)

	err = ioutil.WriteFile(fileName, fileContents, 0777)
	check(err)

	fmt.Printf("Created file with contents:\n%s", fileContents)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	check(err)

	defer r.Body.Close()

	savePostToFile(newPost)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - OK"))
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handling Delete\n"))
}

func editPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handling Put\n"))
}

func handleBlogPosts(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getPost(w, r)
	case "POST":
		createPost(w, r)
	case "DELETE":
		deletePost(w, r)
	case "PUT":
		editPost(w, r)
	}
}

func main() {
	http.HandleFunc("/blogpost", handleBlogPosts)
	fmt.Println("Server listening on port 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
