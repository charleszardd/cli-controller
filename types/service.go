package types

type Service struct {
	ID          	   int    `json:"id"`
	NAME               string `json:"name"`
	DESCRIPTION        string `json:"description"`
}

// type ServiceResponse struct {
// 	Message string `json:"message"`
// }
type ServiceMessageResponse struct {
	Message string `json:"message"`
}