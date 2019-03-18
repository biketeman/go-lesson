package main

import "./blog/firebase"
import "./blog"

func main() {

	// adapter, _ := mongo.New(mongo.MongoDB{
	// 	Address: "mongo://",
	// })

	adapter, _ := firebase.New(firebase.Firebase{
		APIKey:   "dfghiey9834y87y5",
		APIToken: "ihi23y4c98y74983427yc498y37",
	})

	adapter.NewPost(&blog.Post{
		Title: "Mon super post",
	})
	adapter.UpdatePost(&blog.Post{
		Title: "Billy",
	})
	adapter.DeletePost(&blog.Post{
		Title: "Billy",
	})
}
