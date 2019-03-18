package handler 

type ResponsePostEvent struct{
	StatusCode int32 `json:"status"`
	Message string `json: "message"`
	Data *ResponsePostEventData
}

type ResponsePostEventData struct{ 
	Name string
}

type ResponsePostuser struct{
	StatusCode int32 `json:"status`
	Message string `json: "message"`
	Data *ResponsePostUserData

}
type ResponsePostUserData struct{ 
	Username string `json: "name"`
	Email string `json: "email"`
}