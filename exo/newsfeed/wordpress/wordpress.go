package wordpress

import "../../newsfeed"

type WordPress struct {
	URL    string
	APIkey string
}

func New(config Wordpress) (newsfeed.Source, error) {
	return &config, nil
}
func (wp *Wordpress) GetLatestPosts() ([20]*newsfeed.Post, *newsfeed.HTTPerror) {

	return
}
