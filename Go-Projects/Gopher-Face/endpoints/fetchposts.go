package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/go-pool/Go-Data-Structures/socialmedia"
	"github.com/gorilla/mux"
)

/*FetchPostsEndpoint is the end point for fetch posts*/
func FetchPostsEndpoint(w http.ResponseWriter, r *http.Request) {

	// TODO: Implement fetching posts for a given user

	// We are going to create some mock data and send it out in JSON format.

	// We will actually implement this endpoint, when we cover databse persistence later in the course.
	v := mux.Vars(r)

	if v["username"] == "MerJQ" {
		mockPosts := make([]socialmedia.Post, 3)

		post1 := socialmedia.NewPost("MerJQ", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "/images/gopher.png", "", []string{"go", "golang", "programming language"})
		post2 := socialmedia.NewPost("MerJQ", socialmedia.Moods["happy"], "Tour of Go", "Check out the Tour of Go!", "https://tour.golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post3 := socialmedia.NewPost("MerJQ", socialmedia.Moods["hopeful"], "Go Playground", "Check out the Go Playground!", "https://playground.golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)

		json.NewEncoder(w).Encode(mockPosts)
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}
