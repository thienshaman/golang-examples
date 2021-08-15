package entities

type ErrorResponse struct {
	Message string `json:"message"`
}

type FindDataSuccessResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
