package newsfeed

type HTTPerror struct {
	statusCode int32
	Message    string
}

type Newsfeed struct {
	Posts []*Post
}

type Source interface {
	GetLatestPosts() ([20]*Post, HTTPerror)
}
