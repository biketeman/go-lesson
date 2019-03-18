package mongo

import (
	"fmt"

	"../../blog"
)

type MongoDB struct {
	Address  string
	User     string
	Password string
}

func New(db MongoDB) (blog.Blog, error) {

	return &MongoDB{
		Address: db.Address,
	}, nil
}

func (db *MongoDB) NewPost(post *blog.Post) error {

	fmt.Println("New post via MongoDB")

	return nil

}
func (db *MongoDB) UpdatePost(post *blog.Post) error {

	fmt.Println("post updated via MongoDB")

	return nil

}
