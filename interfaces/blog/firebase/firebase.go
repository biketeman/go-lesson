package firebase

import (
	"fmt"

	"../../blog"
)

type Firebase struct {
	APIKey   string
	APIToken string
}

func New(db Firebase) (blog.Blog, error) {

	return &Firebase{
		APIKey:   db.APIKey,
		APIToken: db.APIToken,
	}, nil
}

func (db *Firebase) NewPost(post *blog.Post) error {

	fmt.Println("New post via Firebase")

	return nil

}
func (db *Firebase) UpdatePost(post *blog.Post) error {

	fmt.Println("post updated via Firebase")

	return nil

}
func (db *Firebase) DeletePost(post *blog.Post) error {

	fmt.Println("post was deletd via Firebase")

	return nil

}
